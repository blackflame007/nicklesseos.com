package handlers

import (
	"github.com/blackflame007/nicklesseos.com/app/views"
	"github.com/labstack/echo/v4"
)

type NotFoundHandler struct{}

func (h NotFoundHandler) NotFoundPage(c echo.Context) error {
	return render(c, views.NotFound())
}
