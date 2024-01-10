// nicklesseos.com
package main

import (
	"log/slog"
	"net/http"
	"os"

	"github.com/blackflame007/nicklesseos.com/app/assets"
	"github.com/blackflame007/nicklesseos.com/handlers"
	"github.com/gorilla/sessions"
	"github.com/joho/godotenv"
	"github.com/labstack/echo-contrib/session"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/robfig/cron/v3"
)

func main() {
	// enable .env variables
	err := godotenv.Load()
	if err != nil {
		slog.Warn("Error loading .env file")
	}

	// Use an environment variable for the session key
	sessionKey := os.Getenv("SESSION_KEY")
	if sessionKey == "" {
		slog.Error("Session key is not set. Exiting.")
		return
	}

	// Create Web Server
	app := echo.New()

	// Root level middleware
	app.Use(middleware.Logger())
	app.Use(middleware.Recover())
	app.Use(session.Middleware(sessions.NewCookieStore([]byte(sessionKey))))
	app.Use(func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			cc := &handlers.CustomContext{Context: c}
			return next(cc)
		}
	})

	homeHandler := handlers.HomeHandler{}
	aboutHandler := handlers.AboutHandler{}
	userHandler := handlers.UserHandler{}
	notfoundHandler := handlers.NotFoundHandler{}
	soonHandler := handlers.SoonHandler{}
	// spaceHandler := handlers.NewSpaceManager("https://sfo3.digitaloceanspaces.com", "us-east-1")
	gamePlateformHandler := handlers.GamePlateFormController{}

	// Google OAuth2
	googleHandler := handlers.NewGoogleHandler()

	app.GET("/", homeHandler.IndexPage)

	app.GET("/login", googleHandler.HandleGoogleLogin)
	app.GET("/auth/google/callback", googleHandler.HandleGoogleCallback)
	app.GET("/logout", googleHandler.HandleLogout)

	app.GET("/about", aboutHandler.AboutPage)

	app.GET("/portfolio", soonHandler.SoonPage)

	app.GET("/g", gamePlateformHandler.HandleGamePlateformGallery)

	app.GET("/g/:gameName", gamePlateformHandler.HandleGamePlateformShow)

	// app.GET("/upload", func(c echo.Context) error {

	// 	err := spaceHandler.Upload("app/assets/dist/img/nick_profile.jpg")

	// 	if err != nil {
	// 		return c.String(500, err.Error())
	// 	}

	// 	return c.String(200, "OK")
	// })

	app.GET("/user", userHandler.HandleUserShow)
	// Create a health check endpoint
	app.GET("/health", func(c echo.Context) error {
		return c.String(200, "OK")
	})

	// Create robots.txt endpoint allow all
	app.GET("/robots.txt", func(c echo.Context) error {
		return c.String(200, "User-agent: *\nDisallow:")
	})

	// 404 page
	app.GET("*", notfoundHandler.NotFoundPage)

	// // Serve static files
	// Serve static files with custom headers
	app.GET("/games/*", handlers.HandleGamePlateformStaticFiles)
	// Setup a cron job for cleaning up extracted game files
	c := cron.New()
	_, err = c.AddFunc("@daily", func() {
		handlers.CleanupExtractedFiles("/tmp")
	})
	if err != nil {
		slog.Error("Error setting up cron job for cleanup: ", err)
	}
	c.Start()

	// Serve embedded static files
	app.GET("/dist/*", echo.WrapHandler(http.StripPrefix("/dist/", http.FileServer(assets.CreateFileSystem(false)))))

	app.Logger.Fatal(app.Start(":3000"))
}
