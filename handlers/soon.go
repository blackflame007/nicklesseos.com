package handlers

import (
	"github.com/blackflame007/nicklesseos.com/app/views"
	"github.com/labstack/echo/v4"
)

type SoonHandler struct{}

func (h SoonHandler) SoonPage(c echo.Context) error {
	return render(c, views.Soon())
}
