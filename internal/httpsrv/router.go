package httpsrv

import (
	"os"

	"github.com/bnema/kickstart-echo-htmx/internal/core"
	"github.com/bnema/kickstart-echo-htmx/internal/httpsrv/handler"
	"github.com/bnema/kickstart-echo-htmx/pkg/echo/extra"
	"github.com/gorilla/sessions"
	"github.com/labstack/echo-contrib/session"
	"github.com/labstack/echo/v4"
)

// RegisterRoutes registers the routes for the application
func RegisterRoutes(e *echo.Echo, a *core.App) *echo.Echo {
	// Register middleware
	e.Use(extra.AccessLogMiddleware)
	e.Use(extra.ErrorHandlerMiddleware)
	e.Use(session.Middleware(sessions.NewCookieStore([]byte(os.Getenv("SESSION_SECRET")))))
	e.Use(extra.RestrictSubfoldersMiddleware)

	// Register routes
	bindRootRoute(e, a, "/")
	bindStaticRoute(e, a, "/*")

	return e
}

// bindStaticRoute binds the static route
func bindStaticRoute(e *echo.Echo, a *core.App, path string) {
	e.GET(path, func(c echo.Context) error {
		// set hello world in session
		sess, _ := session.Get("session", c)
		sess.Values["hello"] = "world"
		sess.Save(c.Request(), c.Response())

		return c.String(403, "Forbidden")
	})
}

// bindRootRoute binds the root route
func bindRootRoute(e *echo.Echo, a *core.App, path string) {
	e.GET(path, func(c echo.Context) error {
		return handler.RootPath(c, a)
	})
}
