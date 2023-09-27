package main

import (
	"fmt"
	"log"

	"github.com/bnema/kickstart-echo-htmx/internal/core"
	"github.com/bnema/kickstart-echo-htmx/internal/httpsrv"
	"github.com/labstack/echo/v4"
)

func main() {
	a := core.NewApp()

	e := echo.New()
	e.HideBanner = true
	e.HidePort = true

	e = httpsrv.RegisterRoutes(e, a)

	log.Println("Starting server on port", a.HttpPort, "| Build version", a.BuildVersion)
	if err := e.Start(fmt.Sprintf(":%d", a.HttpPort)); err != nil {
		log.Fatal("Server error:", err)
	}
}
