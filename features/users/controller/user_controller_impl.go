package controller

import (
	"context"
	"fmt"
	"learn-echo/features/users/model/dto"
	"learn-echo/features/users/service"
	ch "learn-echo/pkg/controllerhelper"
	"net/http"

	"github.com/go-playground/mold/v4/modifiers"
	"github.com/labstack/echo/v4"
)

type UserControllerImpl struct {
	UserService service.UserService
	// Validate    *validator.Validate
}

func NewUserController(userService service.UserService) *UserControllerImpl {
	return &UserControllerImpl{
		UserService: userService,
		// Validate:    validate,
	}
}

func (controller *UserControllerImpl) Create(c echo.Context) error {
	var userRequest dto.UserCreateRequest
	conform := modifiers.New()

	errBind := c.Bind(&userRequest)
	if errBind != nil {
		return echo.NewHTTPError(http.StatusBadRequest, errBind.Error())
	}

	err := conform.Struct(context.Background(), &userRequest)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	errVal := c.Validate(userRequest)
	if errVal != nil {
		return echo.NewHTTPError(http.StatusBadRequest, errVal)
	}

	result, err := controller.UserService.Create(userRequest)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "your email or handphone number is already registered")
	}

	return c.JSON(http.StatusCreated, ch.ResponseOkWithData("create data user success", result))
}

func (controller *UserControllerImpl) Login(c echo.Context) error {
	var userRequest dto.UserLoginRequest

	errBind := c.Bind(&userRequest)
	if errBind != nil {
		return echo.NewHTTPError(http.StatusBadRequest, errBind.Error())
	}

	errVal := c.Validate(userRequest)
	if errVal != nil {
		return echo.NewHTTPError(http.StatusBadRequest, errVal)
	}

	result, err := controller.UserService.Login(userRequest)
	if err != nil {
		fmt.Println("err: ", err)
		if err.Error() == "password incorrect" {
			return echo.NewHTTPError(http.StatusUnauthorized, err.Error())
		} else {

			return echo.NewHTTPError(http.StatusNotFound, err.Error())
		}
	}

	return c.JSON(http.StatusOK, ch.ResponseOkWithData("login success", result))
}
