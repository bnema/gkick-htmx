package core

import (
	"io/fs"

	"github.com/bnema/kickstart-echo-htmx/internal/gotmpl"
	"github.com/bnema/kickstart-echo-htmx/internal/webui"
)

const (
	// BuildVersion is the version of the build
	BuildVersion  = "0.0.1"
	HttpPort      = 8181
	mainPath      = "html/"
	fragmentsPath = "html/fragments"
)

type App struct {
	TemplateFS   fs.FS
	PublicFS     fs.FS
	BuildVersion string
	HttpPort     int
	Embed        Embed
}

type AppInterface interface {
	GetTemplateFS() fs.FS
	GetPublicFS() fs.FS
	GetBuildVersion() string
	GetHttpPort() int
	GetEmbed() Embed
}

type Embed struct {
	MainPath      string
	FragmentsPath string
}

func NewApp() *App {
	return &App{
		BuildVersion: BuildVersion,
		HttpPort:     HttpPort,
		TemplateFS:   gotmpl.TemplateFS,
		PublicFS:     webui.PublicFS,
		Embed: Embed{
			MainPath:      mainPath,
			FragmentsPath: fragmentsPath,
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

func (a *App) GetEmbed() Embed {
	return a.Embed
}
