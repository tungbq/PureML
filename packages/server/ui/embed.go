// Package ui handles the PocketBase Admin frontend embedding.
package ui

import (
	"embed"

	"github.com/labstack/echo/v4"
)

//go:embed all:public/build
var buildDir embed.FS

//go:embed all:build
var buildIndex embed.FS

// BuildDirFS contains the embedded build directory files (without the "public/build" prefix)
var BuildDirFS = echo.MustSubFS(buildDir, "public/build")

// BuildIndexFS contains the embedded build directory files (without the "build" prefix)
var BuildIndexFS = echo.MustSubFS(buildIndex, "build")
