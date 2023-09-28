package handler

import (
	"fmt"
	"net/http"
	"os"

	"github.com/bnema/kickstart-echo-htmx/internal/core"
	"github.com/labstack/echo/v4"
)

// StaticRoute serves static files from the embedded filesystem
func StaticRoute(c echo.Context, a *core.App) error {
	// Set the cache-control header based on PROD environment variable
	if os.Getenv("DEV") == "true" {
		fmt.Println("DEV mode: no-cache")
		c.Response().Header().Set("Cache-Control", "no-cache")
	} else {
		c.Response().Header().Set("Cache-Control", "public, max-age=86400")
	}

	// Serve the file from the embedded filesystem
	publicFS := http.FileServer(http.FS(a.PublicFS))
	publicFS.ServeHTTP(c.Response(), c.Request())

	return nil
}
