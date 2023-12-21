package handlers

import (
	"github.com/blackflame007/nicklesseos.com/app/views/user"
	"github.com/labstack/echo/v4"
)

type UserHandler struct{}

func (h UserHandler) HandleUserShow(c echo.Context) error {
	return render(c, user.Show())
}
