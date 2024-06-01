package handlers

import (
	"github.com/a-h/templ"
	"github.com/blackflame007/nicklesseos.com/app/layouts"
	"github.com/labstack/echo-contrib/session"
	"github.com/labstack/echo/v4"
)

func render(c echo.Context, component templ.Component) error {
	// Check if the user is authenticated
	sess, _ := session.Get("session", c)
	isAuthenticated := false
	if auth, ok := sess.Values["authenticated"].(bool); ok && auth {
		isAuthenticated = true
	}

	// Wrap the component in a layout with the authentication status
	layout := layouts.Base(component, isAuthenticated)

	return layout.Render(c.Request().Context(), c.Response())
}

// RenderComponent renders a component with the given context without a layout
func renderComponent(c echo.Context, component templ.Component) error {
	return component.Render(c.Request().Context(), c.Response())
}
