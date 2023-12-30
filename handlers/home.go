package handlers

import (
	"github.com/blackflame007/nicklesseos.com/app/views"
	"github.com/labstack/echo/v4"
)

type HomeHandler struct{}

func (h HomeHandler) IndexPage(c echo.Context) error {
	return render(c, views.Index())
}
