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

// router binds the routes to the echo instance
func router(a *core.App) *echo.Echo {
	e := a.Echo
	e.Debug = true
	e.HideBanner = true
	e.HidePort = true

	e.Use(session.Middleware(sessions.NewCookieStore([]byte(os.Getenv("SESSION_SECRET")))))

	serve.BindRootRoute(a, "/")

	return e
}
