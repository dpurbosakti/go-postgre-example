package controller

import (
	"context"
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
	// errVal := controller.Validate.Struct(userRequest)
	if errVal != nil {
		return echo.NewHTTPError(http.StatusBadRequest, errVal.Error())
	}

	result, err := controller.UserService.Create(userRequest)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "your email or handphone number is already registered")
	}

	return c.JSON(http.StatusCreated, ch.ResponseOkWithData("create data user success", result))
}
