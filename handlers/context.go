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
		// Redirect to login page or return an unauthorized error
		c.Redirect(http.StatusTemporaryRedirect, "/login")
	}

}

func (c *CustomContext) GetData(key string) {

	sess, err := session.Get("session", c)
	if err != nil {
		slog.Error(fmt.Sprintf("Error getting session: %v", err))
	}

	ctx := context.WithValue(c.Request().Context(), "user", sess.Values["user"])

	c.SetRequest(c.Request().WithContext(ctx))
}

func (c *CustomContext) AddGameHeaders() {
	// Set the necessary headers here
	c.Response().Header().Set("Cross-Origin-Embedder-Policy", "require-corp")
	c.Response().Header().Set("Cross-Origin-Opener-Policy", "same-origin")

}
