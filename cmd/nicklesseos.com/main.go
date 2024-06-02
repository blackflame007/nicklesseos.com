// nicklesseos.com
package main

import (
	"fmt"
	"log"
	"log/slog"
	"net/http"
	"os"

	"github.com/blackflame007/nicklesseos.com/app/assets"
	"github.com/blackflame007/nicklesseos.com/handlers"
	service "github.com/blackflame007/nicklesseos.com/services"
	"github.com/gorilla/sessions"
	"github.com/joho/godotenv"
	"github.com/labstack/echo-contrib/session"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
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

	app.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowMethods: []string{http.MethodGet, http.MethodHead, http.MethodPut, http.MethodPatch, http.MethodPost, http.MethodDelete},
	}))
	app.Use(middleware.MethodOverride())

	// spaceHandler := handlers.NewSpaceManager("https://sfo3.digitaloceanspaces.com", "us-east-1")
	dbService, err := service.NewDatabaseService(fmt.Sprintf("%s?authToken=%s", os.Getenv("DB_URL"), os.Getenv("DB_AUTH_TOKEN")))
	if err != nil {
		log.Fatalf("Failed to initialize database service: %s", err)
	}
	userService := service.NewUserService(dbService)

	// Home handler
	homeHandler := handlers.HomeHandler{}
	// About handler
	aboutHandler := handlers.AboutHandler{}
	// User handler
	userHandler := handlers.UserHandler{}
	// 404 handler
	notfoundHandler := handlers.NotFoundHandler{}
	soonHandler := handlers.SoonHandler{}

	// Admin panel handler
	adminHandler := handlers.NewAdminHandler(userService)
	// // Blog handler
	// blogHandler := handlers.BlogHandler{}

	// gamePlateFormController := handlers.NewGamePlateFormController(userService)

	// Google OAuth2
	googleHandler := handlers.NewGoogleHandler(userService)

	app.GET("/", homeHandler.IndexPage)

	app.GET("/login", googleHandler.HandleGoogleLogin)
	app.GET("/auth/google/callback", googleHandler.HandleGoogleCallback)
	app.GET("/logout", googleHandler.HandleLogout)

	app.GET("/about", aboutHandler.AboutPage)

	app.GET("/portfolio", soonHandler.SoonPage)

	// Blog and comments
	// app.GET("/b", blogHandler.HandleBlogIndex)
	// app.GET("/b/:slug", blogHandler.HandleBlogShow)
	// app.POST("/b/:slug/comment", blogHandler.HandleBlogCommentCreate)

	// Admin panel routes
	app.GET("/admin", adminHandler.HandleAdminIndex)
	app.POST("/admin/generate_token", adminHandler.GenerateToken)
	app.GET("/admin/tokens", adminHandler.TokensPage)
	app.DELETE("/admin/revoke_token/:id", adminHandler.RevokeToken)

	// // Admin blog routes
	// app.GET("/admin/blog", adminHandler.HandleAdminBlogIndex)
	// app.GET("/admin/blog/new", adminHandler.HandleAdminBlogNew)
	// app.POST("/admin/blog/new", adminHandler.HandleAdminBlogCreate)
	// app.GET("/admin/blog/edit/:slug", adminHandler.HandleAdminBlogEdit)
	// app.POST("/admin/blog/edit/:slug", adminHandler.HandleAdminBlogUpdate)
	// app.POST("/admin/blog/delete/:slug", adminHandler.HandleAdminBlogDelete)
	// // Admin user routes
	// app.GET("/admin/user", adminHandler.HandleAdminUserIndex)
	// app.GET("/admin/user/edit/:id", adminHandler.HandleAdminUserEdit)

	// API routes
	// app.GET("/api/resource", someHandler, middleware.JWTAuth)

	// User routes
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
	// // Serve static files with custom headers
	// app.GET("/games/*", handlers.HandleGamePlateformStaticFiles)
	// // Setup a cron job for cleaning up extracted game files
	// c := cron.New()
	// _, err = c.AddFunc("@daily", func() {
	// 	handlers.CleanupExtractedFiles("/tmp")
	// })
	// if err != nil {
	// 	slog.Error("Error setting up cron job for cleanup: ", err)
	// }
	// c.Start()

	// Serve embedded static files
	app.GET("/dist/*", echo.WrapHandler(http.StripPrefix("/dist/", http.FileServer(assets.CreateFileSystem(false)))))

	app.Logger.Fatal(app.Start(":3000"))
}
