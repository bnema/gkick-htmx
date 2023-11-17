package core

import (
	"io/fs"
	"log"
	"os"
	"strconv"

	"github.com/bnema/gkick/internal/webui"
	"github.com/bnema/gkick/pkg/http"
	"github.com/bnema/gkick/pkg/render"
	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
)

type App struct {
	BuildVersion string
	Echo         *echo.Echo
	HttpPort     int
	PublicFS     fs.FS
}

func NewApp() *App {
	// Initialize environment variables
	if err := godotenv.Load(); err != nil {
		log.Println("Error loading .env file, continuing with default values")
	}

	// Use default values if environment variables are not set
	httpPort, _ := strconv.Atoi(getEnvDefault("HTTP_PORT", "8080"))

	app := &App{
		HttpPort:     httpPort,
		Echo:         echo.New(),
		PublicFS:     webui.PublicFS,
		BuildVersion: getEnvDefault("BUILD_VERSION", "development"),
	}

	// Initialize template renderer
	renderer, err := render.NewRenderer(app.PublicFS)
	if err != nil {
		log.Fatalf("Failed to create template renderer: %v", err)
	}

	// Pimp the echo instance
	e := app.Echo
	e.Renderer = renderer
	e.StaticFS("/", webui.StaticFS)
	e.Use(http.SetCacheControl)

	return app
}

func getEnvDefault(key, defaultValue string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}
	return defaultValue
}
