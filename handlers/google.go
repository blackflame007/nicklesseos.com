package handlers

import (
	"context"
	"crypto/rand"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io"
	"log/slog"
	"net/http"
	"os"

	"github.com/blackflame007/nicklesseos.com/models"
	service "github.com/blackflame007/nicklesseos.com/services"
	"github.com/labstack/echo-contrib/session"
	"github.com/labstack/echo/v4"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
)

type GoogleHandler struct {
	googleOauthConfig *oauth2.Config
	userService       *service.UserService
}

func NewGoogleHandler(userService *service.UserService) GoogleHandler {
	redirectURL := os.Getenv("GOOGLE_OAUTH_REDIRECT_URL") // Set this in your environment
	if redirectURL == "" {
		redirectURL = "http://localhost:3000/auth/google/callback" // Default for development
	}

	return GoogleHandler{
		googleOauthConfig: &oauth2.Config{
			RedirectURL:  redirectURL,
			ClientID:     os.Getenv("GOOGLE_CLIENT_ID"),
			ClientSecret: os.Getenv("GOOGLE_CLIENT_SECRET"),
			Scopes:       []string{"https://www.googleapis.com/auth/userinfo.email", "https://www.googleapis.com/auth/userinfo.profile", "openid"},
			Endpoint:     google.Endpoint,
		},
	}
}

func (gh GoogleHandler) HandleGoogleLogin(c echo.Context) error {
	// Generate a random state
	b := make([]byte, 32)
	_, err := rand.Read(b)
	if err != nil {
		return c.String(http.StatusInternalServerError, "Failed to generate state: "+err.Error())
	}
	state := base64.StdEncoding.EncodeToString(b)

	// Store the state in session
	sess, _ := session.Get("session", c)
	sess.Values["state"] = state
	sess.Save(c.Request(), c.Response())

	// Include a prompt parameter in the AuthCodeURL method
	url := gh.googleOauthConfig.AuthCodeURL(state, oauth2.AccessTypeOffline, oauth2.ApprovalForce, oauth2.SetAuthURLParam("prompt", "select_account"))
	return c.Redirect(http.StatusTemporaryRedirect, url)
}

func (gh GoogleHandler) HandleGoogleCallback(c echo.Context) error {
	code := c.QueryParam("code")
	token, err := gh.googleOauthConfig.Exchange(context.Background(), code)
	if err != nil {
		return c.String(http.StatusInternalServerError, "Code Exchange failed: "+err.Error())
	}

	response, err := http.Get("https://www.googleapis.com/oauth2/v2/userinfo?access_token=" + token.AccessToken)
	if err != nil {
		return c.String(http.StatusInternalServerError, "Failed to get user info: "+err.Error())
	}
	defer response.Body.Close()

	contents, err := io.ReadAll(response.Body)
	// debug log contents
	fmt.Println("Contents: ", string(contents))

	var userInfo models.UserInfo

	if err != nil {
		return c.String(http.StatusInternalServerError, "Failed to read user info: "+err.Error())
	}

	if err := json.Unmarshal(contents, &userInfo); err != nil {
		return c.String(http.StatusInternalServerError, "Failed to parse user info: "+err.Error())
	}

	gh.userService.CreateUserTable()

	// user := models.User{
	// 	UserInfo: userInfo,
	// }
	// gh.userService.InsertUser(user)

	// Here you can implement logic to create or get a user in your system based on Google userInfo
	sess, _ := session.Get("session", c)
	sess.Values["authenticated"] = true
	sess.Values["user"] = &userInfo.FullName
	sess.Options.MaxAge = 86400 // 24 hours
	sess.Save(c.Request(), c.Response())

	// Redirect or handle the login success
	return c.Redirect(http.StatusTemporaryRedirect, "/")
}

// HandleLogout invalidates a user session and logsout of google locally
func (gh GoogleHandler) HandleLogout(c echo.Context) error {
	sess, err := session.Get("session", c)
	if err != nil {
		slog.Error("Error getting session: ", err)
		return err
	}
	// Clear the session values
	sess.Values["authenticated"] = false
	sess.Values["user"] = ""
	sess.Options.MaxAge = -1 // Expire the session

	err = sess.Save(c.Request(), c.Response())
	if err != nil {
		slog.Error("Error saving session: ", err)
		return err
	}

	return c.Redirect(http.StatusTemporaryRedirect, "/")
}
