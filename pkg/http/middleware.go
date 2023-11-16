package http

import (
	"net/http"
	"os"

	"log/slog"

	"github.com/bnema/gkick/pkg/htmx"
	"github.com/labstack/echo/v4"
)

func SetCacheControl(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		if os.Getenv("DEV") == "true" {
			c.Response().Header().Set("Cache-Control", "no-cache")
		} else {
			c.Response().Header().Set("Cache-Control", "public, max-age=86400")
		}
		return next(c)
	}
}

// AccessLogMiddleware logs access details
func AccessLogMiddleware(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		clientIP := c.RealIP()
		reqMethod := c.Request().Method
		reqPath := c.Request().URL.Path

		err := next(c)

		// Log the access
		slog.Info("access", "client_ip", clientIP, "method", reqMethod, "path", reqPath)
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
			slog.Error("error", err, "client_ip", clientIP, "method", reqMethod, "path", reqPath)

			if he, ok := err.(*echo.HTTPError); ok {
				return c.String(he.Code, he.Message.(string))
			}
			return c.String(http.StatusInternalServerError, "Internal Server Error")
		}
		return nil
	}
}

type Request struct {
	*htmx.Request
}

type Response struct {
	*htmx.Response
}

// EnsureHtmxHeaders checks if the request contains the HTMX headers
func EnsureHtmxHeadersMiddleware() echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(ctx echo.Context) error {
			req, err := htmx.GetRequest(ctx)
			if err != nil {
				return err
			}
			if !req.Enabled {
				return echo.NewHTTPError(http.StatusForbidden, "HTMX headers are missing")
			}
			return next(ctx)
		}
	}
}

// ValidateTriggerNames checks if the request contains a valid trigger name
func ValidateTriggerNamesMiddleware(validTriggers map[string]bool) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(ctx echo.Context) error {
			req, err := htmx.GetRequest(ctx)
			if err != nil {
				return err
			}
			if _, isValid := validTriggers[req.TriggerName]; !isValid {
				return echo.NewHTTPError(http.StatusForbidden, "Invalid trigger name")
			}
			return next(ctx)
		}
	}
}

// ValidateTargetNames checks if the request contains a valid target name
func ValidateTargetNamesMiddleware(validTargets map[string]bool) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(ctx echo.Context) error {
			req, err := htmx.GetRequest(ctx)
			if err != nil {
				return err
			}
			if _, isValid := validTargets[req.Target]; !isValid {
				return echo.NewHTTPError(http.StatusForbidden, "Invalid target name")
			}
			return next(ctx)
		}
	}
}
