// nicklesseos.com
package main

import (
	"github.com/blackflame007/nicklesseos.com/handlers"
	"github.com/labstack/echo/v4"
)

func main() {

	// Create Web Server
	app := echo.New()

	userHandler := handlers.UserHandler{}

	app.GET("/", func(c echo.Context) error {
		return c.String(200, "Hello, World!!!!!!")
	})

	app.GET("/user", userHandler.HandleUserShow)
	// Create a health check endpoint
	app.GET("/health", func(c echo.Context) error {
		return c.String(200, "OK")
	})

	app.Logger.Fatal(app.Start(":3000"))
}
