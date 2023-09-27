package handler

import (
	"github.com/bnema/kickstart-echo-htmx/internal/core"
	"github.com/bnema/kickstart-echo-htmx/pkg/echo/extra"
	"github.com/labstack/echo-contrib/session"
	"github.com/labstack/echo/v4"
)

// RootPath handles the root route to display index.gohtml from the templateFS with the data from strings.yml
func RootPath(c echo.Context, a *core.App) error {
	// Retrieve the session
	sess, err := session.Get("session", c)
	if err != nil {
		return err
	}

	// Retrieve the value from the session
	helloValue := sess.Values["hello"]
	// Create a data map to pass to the renderer
	data := map[string]interface{}{
		"Title":        "Kickstart Echo Htmx",
		"HelloSession": helloValue,
	}

	// Render the template with GenericRenderUtility
	renderedHTML, err := extra.GenericRenderUtility(c, a, a.Embed.MainPath, "index.gohtml", data)
	if err != nil {
		return err
	}

	return c.HTML(200, renderedHTML)
}
