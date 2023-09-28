package extra

import (
	"net/http"

	"log/slog"

	"github.com/labstack/echo/v4"
)

// AccessLogMiddleware logs access details
func AccessLogMiddleware(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		clientIP := c.RealIP()
		reqMethod := c.Request().Method
		reqPath := c.Request().URL.Path

		err := next(c)

		// Log the access
		slog.Info("Access log", "client_ip", clientIP, "method", reqMethod, "path", reqPath)

		return err
	}
}

// ErrorHandlerMiddleware logs the error and client details, then sends an appropriate HTTP response
func ErrorHandlerMiddleware(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		err := next(c)
		if err != nil {
			clientIP := c.RealIP()
			reqMethod := c.Request().Method
			reqPath := c.Request().URL.Path

			// Log the error
			slog.Error("Error log", "error", err, "client_ip", clientIP, "method", reqMethod, "path", reqPath)

			if he, ok := err.(*echo.HTTPError); ok {
				return c.String(he.Code, he.Message.(string))
			}
			return c.String(http.StatusInternalServerError, "Internal Server Error")
		}
		return nil
	}
}

// // RestrictSubfoldersMiddleware restricts access to all subfolders by returning a 403 status
// func RestrictSubfoldersMiddleware(next echo.HandlerFunc) echo.HandlerFunc {
// 	return func(c echo.Context) error {
// 		// folder assets
// 		if strings.HasPrefix(c.Request().URL.Path, "/assets/") {
// 			return c.String(http.StatusForbidden, "Forbidden")
// 		}
// 		// folder css
// 		if strings.HasPrefix(c.Request().URL.Path, "/css/") {
// 			return c.String(http.StatusForbidden, "Forbidden")
// 		}
// 		// folder js
// 		if strings.HasPrefix(c.Request().URL.Path, "/js/") {
// 			return c.String(http.StatusForbidden, "Forbidden")
// 		}
// 		return next(c)
// 	}

// }
