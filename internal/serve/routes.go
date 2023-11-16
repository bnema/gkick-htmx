package serve

import (
	"github.com/bnema/gkick/internal/serve/handler"
	"github.com/bnema/gkick/pkg/core"
	"github.com/labstack/echo-contrib/session"
	"github.com/labstack/echo/v4"
)

// bindRootRoute binds the root route
func BindRootRoute(a *core.App, path string) {
	a.Echo.GET(path, func(c echo.Context) error {
		// set hello world in session
		sess, _ := session.Get("session", c)
		sess.Values["hello"] = "world"
		sess.Save(c.Request(), c.Response())
		return handler.RootPath(c, a)
	})
}
