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
)

func isAuthenticated(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		sess, _ := session.Get("session", c)
		if auth, ok := sess.Values["authenticated"].(bool); !ok || !auth {
			// Redirect to login page or return an unauthorized error
			return c.Redirect(http.StatusTemporaryRedirect, "/login")
		}
		return next(c)
	}
}

func protectedHandler(c echo.Context) error {
	return c.String(http.StatusOK, "protected")
}

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

	homeHandler := handlers.HomeHandler{}
	aboutHandler := handlers.AboutHandler{}
	userHandler := handlers.UserHandler{}
	notfoundHandler := handlers.NotFoundHandler{}
	soonHandler := handlers.SoonHandler{}
	spaceHandler := handlers.NewSpaceManager("https://sfo3.digitaloceanspaces.com", "us-east-1")

	// Google OAuth2
	googleHandler := handlers.NewGoogleHandler()

	app.GET("/", homeHandler.IndexPage)
	app.GET("/protected", protectedHandler, isAuthenticated)

	app.GET("/login", googleHandler.HandleGoogleLogin)
	app.GET("/auth/google/callback", googleHandler.HandleGoogleCallback)
	app.GET("/logout", googleHandler.HandleLogout)

	app.GET("/about", aboutHandler.AboutPage)

	app.GET("/portfolio", soonHandler.SoonPage)

	app.GET("/upload", func(c echo.Context) error {

		err := spaceHandler.Upload("app/assets/dist/img/nick_profile.jpg")

		if err != nil {
			return c.String(500, err.Error())
		}

		return c.String(200, "OK")
	}, isAuthenticated)

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
	// Serve embedded static files
	app.GET("/dist/*", echo.WrapHandler(http.StripPrefix("/dist/", http.FileServer(assets.CreateFileSystem(false)))))

	app.Logger.Fatal(app.Start(":3000"))
}
