package handlers

import (
	"github.com/a-h/templ"
	"github.com/blackflame007/nicklesseos.com/app/layouts"
	"github.com/labstack/echo/v4"
)

func render(c echo.Context, component templ.Component) error {
	// Wrap the component in a layout.
	layout := layouts.Base(component)

	return layout.Render(c.Request().Context(), c.Response())
}
