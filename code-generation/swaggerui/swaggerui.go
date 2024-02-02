// Package swaggerui provides handlers for serving the swagger ui static assets.
package swaggerui

// Package to serve Swagger UI static assets under /swaggerui
// assets were downloaded from swagger-ui-dist v3.52.0
// then swagger-initializer.js was updated with the right path to the
// swagger json file

import (
	"embed"

	"github.com/labstack/echo/v4"
)

var (
	//go:embed assets/*
	dist embed.FS

	//go:embed assets/index.html
	indexHTML embed.FS

	distDirFS     = echo.MustSubFS(dist, "assets")
	distIndexHTML = echo.MustSubFS(indexHTML, "assets")
)

// RegisterHandlers registers the swagger ui static handlers under /swaggerui
func RegisterHandlers(e *echo.Echo) {
	e.FileFS("/swaggerui", "index.html", distIndexHTML)
	e.StaticFS("/swaggerui", distDirFS)
}
