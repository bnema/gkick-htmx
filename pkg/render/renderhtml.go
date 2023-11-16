package render

import (
	"html/template"
	"io"
	"io/fs"
	"path"

	"github.com/labstack/echo/v4"
)

// TemplateRenderer is a custom html/template renderer for Echo framework
type TemplateRenderer struct {
	Templates    *template.Template
	BuildVersion string
	PublicFS     fs.FS
}

// Render renders a template document
func (t *TemplateRenderer) Render(w io.Writer, name string, data interface{}, c echo.Context) error {
	// Add global methods if data is a map
	if viewContext, isMap := data.(map[string]interface{}); isMap {
		viewContext["BuildVersion"] = t.BuildVersion
		// The reverse function generates a URL from a route name and parameters
		viewContext["reverse"] = c.Echo().Reverse
	}

	return t.Templates.ExecuteTemplate(w, name, data)
}

// NewRenderer creates a new TemplateRenderer
func NewRenderer(fs fs.FS) (*TemplateRenderer, error) {
	// Parse the main and fragment templates
	modelsGlob := path.Base("*.html")
	fragmentsGlob := path.Join("fragments", "*.html")
	tmpl, err := template.ParseFS(fs, modelsGlob, fragmentsGlob)
	if err != nil {
		return nil, err
	}

	return &TemplateRenderer{
		Templates: tmpl,
	}, nil
}
