package main

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/gorilla/websocket"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// User model
type User struct {
	ID       uint   `json:"id" gorm:"primaryKey"`
	Username string `json:"username" gorm:"unique"`
	Password string `json:"password"`
}

// Message model
type Message struct {
	ID        uint      `json:"id" gorm:"primaryKey"`
	UserID    uint      `json:"user_id"`
	Username  string    `json:"username"`
	Broadcast string    `json:"broadcast"`
	Content   string    `json:"content"`
	CreatedAt time.Time `json:"created_at"`
}

// Broadcast model
type Broadcast struct {
	Name    string
	Clients map[*websocket.Conn]bool
}

// Global variables
var (
	db         *gorm.DB
	upgrader   = websocket.Upgrader{}
	clients    = make(map[*websocket.Conn]bool)
	broadcasts = make(map[string]*Broadcast)
)

// Main function
func Ex() {
	// Connect to mysql database
	var err error
	dsn := "root:root@tcp(127.0.0.1:3306)/chat?charset=utf8mb4&parseTime=True&loc=Local"
	db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	}

	// Migrate the tables
	db.AutoMigrate(&User{}, &Message{})

	// Create echo instance
	e := echo.New()

	// Use middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	// Define routes
	e.POST("/register", register)
	e.POST("/login", login)
	e.POST("/logout", logout)
	e.GET("/messages/:broadcast", getMessages)
	e.GET("/ws/:broadcast", handleWebsocket)

	// Start server
	e.Logger.Fatal(e.Start(":8080"))
}

// Register handler
func register(c echo.Context) error {
	// Get user data from request body
	user := new(User)
	if err := c.Bind(user); err != nil {
		return err
	}

	// Validate user data
	if user.Username == "" || user.Password == "" {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"message": "Username and password are required",
		})
	}

	// Check if user already exists
	var count int64
	db.Model(&User{}).Where("username = ?", user.Username).Count(&count)
	if count > 0 {
		return c.JSON(http.StatusConflict, map[string]string{
			"message": "Username already taken",
		})
	}

	// Create user in database
	db.Create(user)

	// Return success response
	return c.JSON(http.StatusCreated, map[string]string{
		"message": "User registered successfully",
	})
}

// Login handler
func login(c echo.Context) error {
	// Get user data from request body
	user := new(User)
	if err := c.Bind(user); err != nil {
		return err
	}

	// Validate user data
	if user.Username == "" || user.Password == "" {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"message": "Username and password are required",
		})
	}

	// Check if user exists and password matches
	var dbUser User
	db.Where("username = ?", user.Username).First(&dbUser)
	if dbUser.ID == 0 || dbUser.Password != user.Password {
		return c.JSON(http.StatusUnauthorized, map[string]string{
			"message": "Invalid username or password",
		})
	}

	// Return success response
	return c.JSON(http.StatusOK, map[string]string{
		"message": "User logged in successfully",
	})
}

// Logout handler
func logout(c echo.Context) error {
	// Get websocket connection from request context
	conn := c.Get("conn").(*websocket.Conn)

	// Remove connection from clients map
	delete(clients, conn)

	// Close connection
	conn.Close()

	// Return success response
	return c.JSON(http.StatusOK, map[string]string{
		"message": "User logged out successfully",
	})
}

// Get messages handler
func getMessages(c echo.Context) error {
	// Get broadcast name from request parameter
	broadcast := c.Param("broadcast")

	// Validate broadcast name
	if broadcast == "" {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"message": "Broadcast name is required",
		})
	}

	// Get messages from database by broadcast name
	var messages []Message
	db.Where("broadcast = ?", broadcast).Order("created_at asc").Find(&messages)

	// Return success response
	return c.JSON(http.StatusOK, messages)
}

// Handle websocket handler
func handleWebsocket(c echo.Context) error {
	// Get broadcast name from request parameter
	broadcast := c.Param("broadcast")

	// Validate broadcast name
	if broadcast == "" {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"message": "Broadcast name is required",
		})
	}

	// Upgrade http connection to websocket connection
	conn, err := upgrader.Upgrade(c.Response(), c.Request(), nil)
	if err != nil {
		log.Println(err)
		return err
	}

	// Add connection to clients map
	clients[conn] = true

	// Join broadcast
	joinBroadcast(conn, broadcast)

	// Handle websocket events
	go onConnect(conn, broadcast)
	go onMessage(conn, broadcast)
	go onClose(conn, broadcast)

	// Return success response
	return c.JSON(http.StatusOK, map[string]string{
		"message": "Websocket connection established",
	})
}

// On connect event
func onConnect(conn *websocket.Conn, broadcast string) {
	// Send welcome message to client
	conn.WriteJSON(map[string]string{
		"message": fmt.Sprintf("Welcome to %s broadcast", broadcast),
	})
}

// On message event
func onMessage(conn *websocket.Conn, broadcast string) {
	// Loop to read messages from client
	for {
		// Read message from client
		var msg Message
		err := conn.ReadJSON(&msg)
		if err != nil {
			log.Println(err)
			break
		}

		// Save message to database
		db.Create(&msg)

		// Send message to broadcast
		sendMessage(conn, broadcast, msg)
	}
}

// On close event
func onClose(conn *websocket.Conn, broadcast string) {
	// Remove connection from clients map
	delete(clients, conn)

	// Leave broadcast
	leaveBroadcast(conn, broadcast)

	// Close connection
	conn.Close()
}

// Create broadcast function
func createBroadcast(name string) {
	// Create a new broadcast with the given name
	broadcast := &Broadcast{
		Name:    name,
		Clients: make(map[*websocket.Conn]bool),
	}

	// Add broadcast to broadcasts map
	broadcasts[name] = broadcast
}

// Join broadcast function
func joinBroadcast(conn *websocket.Conn, name string) {
	// Check if broadcast exists
	broadcast, ok := broadcasts[name]
	if !ok {
		// If not, create a new broadcast
		createBroadcast(name)
		broadcast = broadcasts[name]
	}

	// Add connection to broadcast clients map
	broadcast.Clients[conn] = true
}

// Leave broadcast function
func leaveBroadcast(conn *websocket.Conn, name string) {
	// Check if broadcast exists
	broadcast, ok := broadcasts[name]
	if !ok {
		// If not, do nothing
		return
	}

	// Remove connection from broadcast clients map
	delete(broadcast.Clients, conn)
}

// Send message function
func sendMessage(conn *websocket.Conn, name string, msg Message) {
	// Check if broadcast exists
	broadcast, ok := broadcasts[name]
	if !ok {
		// If not, do nothing
		return
	}

	// Loop through broadcast clients
	for client := range broadcast.Clients {
		// Check if client is not the sender
		if client != conn {
			// Send message to client
			err := client.WriteJSON(msg)
			if err != nil {
				log.Println(err)
				continue
			}
		}
	}
}
