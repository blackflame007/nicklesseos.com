// nicklesseos.com
package main

import (
	"net/http"

	"github.com/blackflame007/nicklesseos.com/app/assets"
	"github.com/blackflame007/nicklesseos.com/handlers"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {

	// Create Web Server
	app := echo.New()

	// Root level middleware
	app.Use(middleware.Logger())
	app.Use(middleware.Recover())

	homeHandler := handlers.HomeHandler{}
	aboutHandler := handlers.AboutHandler{}
	userHandler := handlers.UserHandler{}
	notfoundHandler := handlers.NotFoundHandler{}
	soonHandler := handlers.SoonHandler{}

	app.GET("/", homeHandler.IndexPage)

	app.GET("/about", aboutHandler.AboutPage)

	app.GET("/portfolio", soonHandler.SoonPage)

	app.GET("/user", userHandler.HandleUserShow)
	// Create a health check endpoint
	app.GET("/health", func(c echo.Context) error {
		return c.String(200, "OK")
	})

	// 404 page
	app.GET("*", notfoundHandler.NotFoundPage)

	// // Serve static files
	// Serve embedded static files
	app.GET("/dist/*", echo.WrapHandler(http.StripPrefix("/dist/", http.FileServer(assets.CreateFileSystem(false)))))

	app.Logger.Fatal(app.Start(":3000"))
}
