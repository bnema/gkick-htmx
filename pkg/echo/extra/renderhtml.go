package extra

import (
	"github.com/bnema/kickstart-echo-htmx/internal/core"
	"github.com/bnema/kickstart-echo-htmx/pkg/gorender"
	"github.com/labstack/echo/v4"
)

// GenericRenderUtility renders a common HTML template with the provided data.
func GenericRenderUtility(c echo.Context, a *core.App, path string, templateName string, customData map[string]interface{}) (string, error) {

	// Navigate inside the fs.FS to get the template
	rendererData, err := gorender.GetHTMLRenderer(path, templateName, a.TemplateFS, a)
	if err != nil {
		return "", err
	}

	// Create a data map to pass to the renderer
	data := map[string]interface{}{
		"BuldVersion": a.BuildVersion,
	}

	// Merge customData into data
	for k, v := range customData {
		data[k] = v
	}

	renderedHTML, err := rendererData.Render(data, a)
	if err != nil {
		return "", err
	}

	return renderedHTML, nil
}
