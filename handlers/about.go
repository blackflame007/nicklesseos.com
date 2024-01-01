package handlers

import (
	"github.com/blackflame007/nicklesseos.com/app/views"
	"github.com/labstack/echo/v4"
)

type AboutHandler struct{}

func (h AboutHandler) AboutPage(c echo.Context) error {
	return render(c, views.About())
}
