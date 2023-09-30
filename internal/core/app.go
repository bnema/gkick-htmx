package core

import (
	"fmt"
	"io/fs"
	"log"
	"os"
	"strconv"

	"github.com/bnema/kickstart-echo-htmx/internal/gotmpl"
	"github.com/bnema/kickstart-echo-htmx/internal/webui"
	"github.com/bnema/kickstart-echo-htmx/pkg/gorender"
	"github.com/joho/godotenv"
)

type App struct {
	BuildVersion           string
	HttpPort               int
	*gorender.RenderConfig // Embed the extra.RenderConfig struct
}

type AppInterface interface {
	GetTemplateFS() fs.FS
	GetPublicFS() fs.FS
	GetBuildVersion() string
	GetHttpPort() int
}

func InitEnv() {
	err := godotenv.Load()
	if err != nil {
		fmt.Println("Error loading .env file")
	}
}

func NewApp() *App {

	InitEnv()
	httpPortStr := os.Getenv("HTTP_PORT")
	httpPort, err := strconv.Atoi(httpPortStr)
	if err != nil {
		log.Println("HTTP_PORT not set, using default 8080")
		httpPort = 8080
	}

	return &App{
		HttpPort: httpPort,
		RenderConfig: &gorender.RenderConfig{
			TemplateFS: gotmpl.TemplateFS,
			PublicFS:   webui.PublicFS,
		},
	}
}

func (a *App) GetTemplateFS() fs.FS {
	return a.TemplateFS
}

func (a *App) GetPublicFS() fs.FS {
	return a.PublicFS
}

func (a *App) GetBuildVersion() string {
	return a.BuildVersion
}

func (a *App) GetHttpPort() int {
	return a.HttpPort
}
