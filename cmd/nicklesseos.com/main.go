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

	app.GET("/user", userHandler.HandleUserShow)
	app.Logger.Fatal(app.Start(":3000"))
}
