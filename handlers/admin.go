package handlers

import (
	// "github.com/blackflame007/nicklesseos.com/app/views/admin"

	"database/sql"
	"fmt"
	"net/http"
	"strconv"
	"time"

	"github.com/blackflame007/nicklesseos.com/app/components"
	"github.com/blackflame007/nicklesseos.com/app/views/admin"
	service "github.com/blackflame007/nicklesseos.com/services"
	"github.com/blackflame007/nicklesseos.com/utils"
	"github.com/labstack/echo-contrib/session"
	"github.com/labstack/echo/v4"
)

type AdminHandler struct {
	userService *service.UserService
}

func NewAdminHandler(userService *service.UserService) *AdminHandler {
	return &AdminHandler{
		userService: userService,
	}
}

func (h AdminHandler) HandleAdminIndex(c echo.Context) error {
	cc := c.(*CustomContext)
	cc.IsAuthenticated()
	cc.IsAdmin()
	return render(cc, admin.Dashboard())
}

func (h AdminHandler) GenerateToken(c echo.Context) error {
	cc := c.(*CustomContext)
	cc.IsAuthenticated()
	cc.IsAdmin()

	// Get email from session
	// Check if the user is authenticated
	sess, _ := session.Get("session", c)
	email, ok := sess.Values["email"].(string)
	if !ok {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "email is required"})
	}
	// Find the user by email
	user, err := h.userService.FindUserByEmail(email)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "User not found"})
	}

	err = h.userService.CreateTokenTable()
	if err != nil {
		fmt.Println(err)
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Failed to create token table"})
	}

	name := c.FormValue("name")

	var expirationTime *time.Time
	if revokedAtStr := c.FormValue("revoked_at"); revokedAtStr != "" {
		parsedTime, err := time.Parse("2006-01-02", revokedAtStr)
		if err != nil {
			return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid date format"})
		}
		expirationTime = &parsedTime
	} else {
		expirationTime = nil // Token will never expire
	}

	token, err := utils.GenerateToken(email, expirationTime)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Failed to generate token"})
	}

	revokedAt := sql.NullTime{
		Time:  time.Time{},
		Valid: false,
	}
	if expirationTime != nil {
		revokedAt = sql.NullTime{
			Time:  *expirationTime,
			Valid: true,
		}
	}

	err = h.userService.SaveToken(user.ID, token, name, revokedAt)
	if err != nil {
		fmt.Println(err)
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Failed to save token"})
	}

	return renderComponent(c, components.Alert(token, "info"))
}

func (h AdminHandler) TokensPage(c echo.Context) error {
	cc := c.(*CustomContext)
	cc.IsAuthenticated()
	cc.IsAdmin()

	// Get email from session
	sess, err := session.Get("session", c)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Failed to get session"})
	}
	email := sess.Values["email"].(string)

	if email == "" {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "email is required"})
	}

	// Find the user by email
	user, err := h.userService.FindUserByEmail(email)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "User not found"})
	}

	tokens, err := h.userService.GetTokensByUserID(user.ID)
	if err != nil {
		fmt.Println(err)
		// return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Failed to retrieve tokens"})
	}

	return render(cc, admin.Tokens(tokens))
}

// TokenFormComponent
func (h AdminHandler) TokenFormComponent(c echo.Context) error {
	cc := c.(*CustomContext)
	cc.IsAuthenticated()
	cc.IsAdmin()

	// Check if the form is currently shown or hidden
	formShown := c.QueryParam("formShown") == "true"
	if formShown {
		return c.HTML(http.StatusOK, `<div id="formToggleState" style="display:none;"></div>`) // Hide the form
	} else {
		return renderComponent(c, admin.TokenForm()) // Show the form
	}
	// return renderComponent(c, admin.TokenForm())
}

func (h AdminHandler) RevokeToken(c echo.Context) error {
	cc := c.(*CustomContext)
	cc.IsAuthenticated()
	cc.IsAdmin()

	tokenIDStr := c.Param("id")
	tokenID, err := strconv.Atoi(tokenIDStr)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid token ID"})
	}

	err = h.userService.RevokeToken(tokenID)
	if err != nil {
		fmt.Println(err)
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Failed to revoke token"})
	}

	return renderComponent(c, components.Toast(fmt.Sprintf("Token ID: %d has been deleted", tokenID), "warning")) // Show the form
}
