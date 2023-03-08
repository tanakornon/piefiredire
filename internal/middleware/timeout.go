package middleware

import (
	"fmt"
	"time"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func NewTimeoutMiddleware(duration int) echo.MiddlewareFunc {
	timeout := time.Duration(duration) * time.Second
	message := fmt.Sprintf("Can't fulfilled the request with in %d seconds.", duration)

	return middleware.TimeoutWithConfig(middleware.TimeoutConfig{
		Skipper:      middleware.DefaultSkipper,
		ErrorMessage: message,
		Timeout:      timeout,
	})
}
