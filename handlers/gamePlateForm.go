package handlers

import (
	"archive/zip"
	"encoding/json"
	"fmt"
	"io"
	"log/slog"
	"net/http"
	"os"
	"path"
	"path/filepath"
	"strings"

	"github.com/blackflame007/nicklesseos.com/app/views/gamePlateform"
	"github.com/blackflame007/nicklesseos.com/models"
	service "github.com/blackflame007/nicklesseos.com/services"
	"github.com/labstack/echo/v4"
)

// GamePlateFormController handles all the platform methods
type GamePlateFormController struct {
	userService *service.UserService
}

func NewGamePlateFormController(userService *service.UserService) *GamePlateFormController {
	return &GamePlateFormController{
		userService: userService,
	}
}

// HandleGamePlateformGallery is the method to control code for the gallery
func (h GamePlateFormController) HandleGamePlateformGallery(c echo.Context) error {
	cc := c.(*CustomContext)
	cc.IsAuthenticated()
	cc.GetData("user")

	spaceHandler := NewSpaceManager("https://sfo3.digitaloceanspaces.com", "us-east-1")

	files, err := spaceHandler.GetFiles("games")
	if err != nil {
		slog.Error(err.Error())
		return err
	}

	// get only keys with zip ext and split name
	var gamesInfo []models.GameContent // Slice to store game information

	for _, file := range files {
		if strings.Contains(file, ".zip") {
			baseName := strings.Split(file, ".")[0]
			gameName := strings.Split(baseName, "/")[1]

			jsonKey := fmt.Sprintf("games/%s.json", gameName)
			if contains(files, jsonKey) {
				content, err := spaceHandler.GetFileContents(jsonKey)
				if err != nil {
					slog.Error("Error getting contents of file: ", jsonKey, " Error: ", err)
					continue
				}

				var gameInfo models.GameContent
				err = json.Unmarshal(content, &gameInfo) // Unmarshal into a new GameContent object
				if err != nil {
					slog.Error("Error unmarshalling json: ", jsonKey, " Error: ", err)
					continue
				}

				// Set the URL field for the game
				gameInfo.URL = fmt.Sprintf("/g/%s", gameName)

				gamesInfo = append(gamesInfo, gameInfo) // Add the gameInfo to the slice
			}
		}
	}

	return render(cc, gamePlateform.Gallery(gamesInfo)) // Pass the slice of GameContent objects to the gallery view
}

// Utility function to check if a slice contains a string
func contains(slice []string, str string) bool {
	for _, v := range slice {
		if v == str {
			return true
		}
	}
	return false
}

// HandleGamePlateformShow is the method to handle an individual game
func (h GamePlateFormController) HandleGamePlateformShow(c echo.Context) error {
	cc := c.(*CustomContext)
	cc.IsAuthenticated()
	cc.GetData("user")
	cc.AddGameHeaders()

	// Get the URL param for the game name
	gameName := c.Param("gameName")

	spaceHandler := NewSpaceManager("https://sfo3.digitaloceanspaces.com", "us-east-1")

	// Construct the paths
	zipPath := fmt.Sprintf("/games/%s.zip", gameName)
	localZipPath := fmt.Sprintf("/tmp/%s.zip", gameName)    // Temporary local path
	extractedFolderPath := fmt.Sprintf("/tmp/%s", gameName) // Extracted files path

	// Download the zip file
	err := spaceHandler.DownloadFile(zipPath, localZipPath)
	if err != nil {
		return err // Handle error
	}

	// Extract the zip file
	err = extractZip(localZipPath, extractedFolderPath)
	if err != nil {
		return err // Handle error
	}

	leaderboard, err := h.userService.GetLeaderboard()
	if err != nil {
		slog.Error("Error getting leaderboard: ", err)
		return err
	}
	fmt.Println(leaderboard)

	// Pass the extracted game path to the render function
	return render(cc, gamePlateform.Show(gameName))
}

// extractZip extracts a zip file to a specified destination
func extractZip(src, dest string) error {
	r, err := zip.OpenReader(src)
	if err != nil {
		return err
	}
	defer r.Close()

	os.MkdirAll(dest, 0755)

	for _, f := range r.File {
		// Log each file extraction
		rc, err := f.Open()
		if err != nil {
			return err
		}
		defer rc.Close()

		path := filepath.Join(dest, f.Name)

		if f.FileInfo().IsDir() {
			os.MkdirAll(path, f.Mode())
		} else {
			var fdir string
			if lastIndex := strings.LastIndex(path, string(os.PathSeparator)); lastIndex > -1 {
				fdir = path[:lastIndex]
			}

			err = os.MkdirAll(fdir, 0755)
			if err != nil {
				return err
			}
			f, err := os.OpenFile(path, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, f.Mode())
			if err != nil {
				return err
			}
			defer f.Close()

			_, err = io.Copy(f, rc)
			if err != nil {
				return err
			}
		}

	}
	return nil
}

// CleanupExtractedFiles removes extracted game files from a directory
func CleanupExtractedFiles(dir string) {
	// Example: Remove all files and directories in 'dir'
	d, err := os.Open(dir)
	if err != nil {
		slog.Error("Error opening directory for cleanup: ", err)
		return
	}
	defer d.Close()

	names, err := d.Readdirnames(-1)
	if err != nil {
		slog.Error("Error reading directory names for cleanup: ", err)
		return
	}

	for _, name := range names {
		err = os.RemoveAll(path.Join(dir, name))
		if err != nil {
			slog.Error("Error removing file/directory: ", name, " Error: ", err)
		}
	}
}

func HandleGamePlateformStaticFiles(c echo.Context) error {
	cc := c.(*CustomContext)
	cc.IsAuthenticated()
	cc.GetData("user")
	cc.AddGameHeaders()

	// Manually handle the file serving
	fileServer := http.StripPrefix("/games/", http.FileServer(http.Dir("/tmp")))
	fileServer.ServeHTTP(c.Response().Writer, c.Request())

	return nil
}
