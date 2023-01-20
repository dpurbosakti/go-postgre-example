package middlewares

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func RemoveTrailingSlash() echo.MiddlewareFunc {
	return middleware.RemoveTrailingSlash()
}

func Slash(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		RemoveTrailingSlash()
		return next(c)
	}
}
