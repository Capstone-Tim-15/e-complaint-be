package main

import (
	"ecomplaint/config"
	"ecomplaint/routes"
	"net/http"

	"github.com/go-playground/validator"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	app := echo.New()
	validate := validator.New()
	DB := config.ConnectDB()
	AI := config.ConnectOpenAI()

	app.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Welcome to RESTful API Services")
	})

	routes.UserRoutes(app, DB, validate)
	routes.AdminRoutes(app, DB, validate)
	routes.OTPRoutes(app, DB, validate)
	routes.ComplaintRoutes(app, DB, validate)
	routes.CommentRoutes(app, DB, validate)
	routes.AIRoutes(app, AI)

	routes.FAQRoutes(app, DB, validate)
	routes.CategoryRoutes(app, DB, validate)

	routes.NewsRoutes(app, DB, validate)
	routes.LikeRoutes(app, DB, validate)
	routes.FeedbackRoutes(app, DB, validate)

	app.Pre(middleware.RemoveTrailingSlash())
	app.Use(middleware.CORS())
	app.Use(middleware.LoggerWithConfig(
		middleware.LoggerConfig{
			Format: "ip: ${host} | method: ${method} | uri: ${uri} | status: ${status}\n",
		},
	))

	app.Logger.Fatal(app.Start(":8000"))

}
