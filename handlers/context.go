package handlers

import (
	"context"
	"fmt"
	"log/slog"
	"net/http"

	"github.com/labstack/echo-contrib/session"
	"github.com/labstack/echo/v4"
)

// CustomContext is a custom context we controll and add methods to
type CustomContext struct {
	echo.Context
}

// IsAuthenticated gets the saved session and checks if the user is Authenticated
func (c *CustomContext) IsAuthenticated() {

	sess, err := session.Get("session", c)
	if err != nil {
		slog.Error(fmt.Sprintf("Error getting session: %v", err))
	}

	if auth, ok := sess.Values["authenticated"].(bool); !ok || !auth {
		// Store the original URL
		sess.Values["original_url"] = c.Request().URL.String()
		sess.Save(c.Request(), c.Response())

		// Redirect to previous page or return an unauthorized error
		err := c.Redirect(http.StatusTemporaryRedirect, "/login")
		if err != nil {
			slog.Error(fmt.Sprintf("Error redirecting to login: %v", err))
		}
	}

}

// IsAdmin checks if the user is an admin
// check if rank is 100
func (c *CustomContext) IsAdmin() {
	sess, err := session.Get("session", c)
	if err != nil {
		slog.Error(fmt.Sprintf("Error getting session: %v", err))
	}

	if rank, ok := sess.Values["rank"].(int); !ok || rank != 100 {
		// Redirect to login page
		err := c.Redirect(http.StatusTemporaryRedirect, "/login")
		if err != nil {
			slog.Error(fmt.Sprintf("Error redirecting to login: %v", err))
		}
	}
}

func (c *CustomContext) GetData(key string) {

	sess, err := session.Get("session", c)
	if err != nil {
		slog.Error(fmt.Sprintf("Error getting session: %v", err))
	}

	// ctx := context.WithValue(c.Request().Context(), "user", sess.Values["user"])
	ctx := context.WithValue(c.Request().Context(), key, sess.Values[key])

	c.SetRequest(c.Request().WithContext(ctx))
}

func (c *CustomContext) AddGameHeaders() {
	// Set the necessary headers here
	c.Response().Header().Set("Cross-Origin-Embedder-Policy", "require-corp")
	c.Response().Header().Set("Cross-Origin-Opener-Policy", "same-origin")

}
