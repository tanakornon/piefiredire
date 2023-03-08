package middleware

import "github.com/labstack/echo/v4"

func NewCustomMiddleware() echo.MiddlewareFunc {
	return customMiddleware
}

func customMiddleware(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		if err := next(c); err != nil {
			c.Error(err)
		}
		return nil
	}
}
