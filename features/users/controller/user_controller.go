package controller

import "github.com/labstack/echo/v4"

type UserController interface {
	Create(c echo.Context) error
	Login(c echo.Context) error
	GetDetail(c echo.Context) error
	GetList(c echo.Context) error
	Delete(c echo.Context) error
	Update(c echo.Context) error
	Verify(c echo.Context) error
	RefreshVerCode(c echo.Context) error
}
