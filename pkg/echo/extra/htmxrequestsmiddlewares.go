package extra

import (
	"net/http"

	"github.com/bnema/kickstart-echo-htmx/pkg/htmx"
	"github.com/labstack/echo/v4"
)

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
