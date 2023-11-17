package main

import (
	"fmt"
	"log"
	"os"

	"github.com/bnema/gkick/internal/serve"
	"github.com/bnema/gkick/pkg/core"
	"github.com/gorilla/sessions"
	"github.com/labstack/echo-contrib/session"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	// Create a new app
	a := core.NewApp()

	// Bind the routes
	router(a)

	log.Println("Starting server on port", a.HttpPort, "| Build version", a.BuildVersion)
	if err := a.Echo.Start(fmt.Sprintf(":%d", a.HttpPort)); err != nil {
		log.Fatal("Server error:", err)
	}
}

// Bindings are created in the serve package and listed here
func router(a *core.App) *echo.Echo {

	serve.BindRootRoute(a, "/") // Return the index page

	// Quick hello world example for htmx get request
	a.Echo.GET("/api/hello",
		func(c echo.Context) error {
			return c.String(200, "Hello, World!")
		},
	)

	// Static files are automatically served from the public/assets directory to / (root) (Example /assets/css/style.css is served to /css/style.css)

	return echoConfig(a)
}

// Example of how to configure echo
func echoConfig(a *core.App) *echo.Echo {
	e := a.Echo
	e.Debug = true
	e.HideBanner = true
	e.HidePort = true
	e.Use(middleware.Recover())
	e.Use(middleware.Secure())
	e.Use(middleware.CORS())
	e.Use(middleware.Gzip())
	e.Use(middleware.Logger())

	e.Use(session.Middleware(sessions.NewCookieStore([]byte(os.Getenv("SESSION_SECRET")))))

	return e
}
