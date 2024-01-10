package handlers

import (
	"github.com/blackflame007/nicklesseos.com/app/views/user"
	"github.com/labstack/echo/v4"
)

type UserHandler struct{}

func (h UserHandler) HandleUserShow(c echo.Context) error {
	cc := c.(*CustomContext)
	cc.IsAuthenticated()
	cc.GetData("user")
	return render(cc, user.Show())
}
