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

	app.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Welcome to RESTful API Services")
	})

	routes.UserRoutes(app, DB, validate)
	routes.AdminRoutes(app, DB, validate)
	routes.OTPRoutes(app, DB, validate)

	app.Pre(middleware.RemoveTrailingSlash())
	app.Use(middleware.CORS())
	app.Use(middleware.LoggerWithConfig(
		middleware.LoggerConfig{
			Format: "method=${method}, uri=${uri}, status=${status}\n",
		},
	))

	app.Logger.Fatal(app.Start(":8000"))

}
