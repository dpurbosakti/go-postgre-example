package controller

import "github.com/labstack/echo/v4"

type AccountController interface {
	Create(c echo.Context) error
}
