package controller

import "github.com/labstack/echo/v4"

type UserController interface {
	Create(c echo.Context) error
}
