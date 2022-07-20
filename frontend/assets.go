package frontend

import (
	"embed"
	"io/fs"
)

//go:embed dist/*
var assets embed.FS

func Value() fs.FS {
	assetsFs, err := fs.Sub(assets, "dist")
	if err != nil {
		panic(err)
	}
	return assetsFs
}
