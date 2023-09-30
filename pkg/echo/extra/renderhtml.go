package extra

import (
	"github.com/bnema/kickstart-echo-htmx/pkg/gorender"
	"github.com/labstack/echo/v4"
)

// GenericRenderUtility renders a common HTML template with the provided data.
func GenericRenderUtility(c echo.Context, templateName string, customData map[string]interface{}, r *gorender.RenderConfig) (string, error) {

	// Navigate inside the fs.FS to get the template
	rendererData, err := gorender.GetHTMLRenderer(templateName, r.TemplateFS)
	if err != nil {
		return "", err
	}

	// Create a data map to pass to the renderer
	data := map[string]interface{}{
		"BuldVersion": r.BuildVersion,
	}

	// Merge customData into data
	for k, v := range customData {
		data[k] = v
	}

	renderedHTML, err := rendererData.Render(data)
	if err != nil {
		return "", err
	}

	return renderedHTML, nil
}
