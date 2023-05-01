package ui

import (
	"embed"
	"io/fs"
)

const distDir = "assets/ui/dist"

//go:embed assets/ui/dist
var distFS embed.FS

func GetFS() fs.FS {
	uiFS, err := fs.Sub(distFS, distDir)
	if err != nil {
		panic(err)
	}

	return uiFS
}
