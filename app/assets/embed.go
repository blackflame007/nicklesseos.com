package assets

import (
	"embed"
	"io/fs"
	"net/http"
)

//go:embed all:dist
var embeddedFiles embed.FS

// createFileSystem helps to create a file system for the embedded files.
func CreateFileSystem(useOS bool) http.FileSystem {
	if useOS {
		return http.FS(embeddedFiles)
	}
	fsys, err := fs.Sub(embeddedFiles, "dist")
	if err != nil {
		panic(err)
	}
	return http.FS(fsys)
}
