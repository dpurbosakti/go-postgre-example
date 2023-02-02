package controller

import (
	"context"
	"learn-echo/features/users/model/dto"
	"learn-echo/features/users/service"
	"learn-echo/middlewares"
	ch "learn-echo/pkg/controllerhelper"
	"learn-echo/pkg/pagination"
	"net/http"
	"strconv"

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
		return echo.NewHTTPError(http.StatusBadRequest, errVal.Error())
	}

	result, err := controller.UserService.Create(userRequest)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
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
		return echo.NewHTTPError(http.StatusBadRequest, errVal.Error())
	}

	result, err := controller.UserService.Login(userRequest)
	if err != nil {
		if err.Error() == "password incorrect" {
			return echo.NewHTTPError(http.StatusUnauthorized, err.Error())
		} else {

			return echo.NewHTTPError(http.StatusNotFound, err.Error())
		}
	}

	return c.JSON(http.StatusOK, ch.ResponseOkWithData("login success", result))
}

func (controller *UserControllerImpl) GetDetail(c echo.Context) error {
	dataToken, _ := middlewares.ExtractToken(c)
	userId := int(dataToken.Id)
	result, err := controller.UserService.GetDetail(userId)
	if err != nil {
		return echo.NewHTTPError(http.StatusNotFound, err)
	}

	return c.JSON(http.StatusOK, ch.ResponseOkWithData("get data user success", result))
}

func (controller *UserControllerImpl) GetList(c echo.Context) error {
	var page pagination.Pagination
	limitInt, _ := strconv.Atoi(c.QueryParam("limit"))
	pageInt, _ := strconv.Atoi(c.QueryParam("page"))
	page.Limit = limitInt
	page.Page = pageInt
	page.Sort = c.QueryParam("sort")

	result, err := controller.UserService.GetList(page)
	if err != nil {
		return echo.NewHTTPError(http.StatusNotFound, err)
	}

	if result.TotalRows == 0 {
		return echo.NewHTTPError(http.StatusNotFound, "data not found")
	}

	return c.JSON(http.StatusOK, ch.ResponseOkWithData("get data users success", result))
}

func (controller *UserControllerImpl) Delete(c echo.Context) error {
	dataToken, _ := middlewares.ExtractToken(c)
	userId := int(dataToken.Id)
	err := controller.UserService.Delete(userId)
	if err != nil {
		return echo.NewHTTPError(http.StatusNotFound, err)
	}

	return c.JSON(http.StatusOK, ch.ResponseOkNoData("delete data user success"))
}

func (controller *UserControllerImpl) Update(c echo.Context) error {
	var userRequest dto.UserUpdateRequest
	conform := modifiers.New()

	errBind := c.Bind(&userRequest)
	if errBind != nil {
		return echo.NewHTTPError(http.StatusBadRequest, errBind.Error())
	}

	err := conform.Struct(context.Background(), &userRequest)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	dataToken, _ := middlewares.ExtractToken(c)
	userId := int(dataToken.Id)
	result, err := controller.UserService.Update(userRequest, userId)
	if err != nil {
		return echo.NewHTTPError(http.StatusNotFound, err)
	}

	return c.JSON(http.StatusOK, ch.ResponseOkWithData("update data users success", result))
}

func (controller *UserControllerImpl) Verify(c echo.Context) error {
	var userRequest dto.UserVerifyRequest

	errBind := c.Bind(&userRequest)
	if errBind != nil {
		return echo.NewHTTPError(http.StatusBadRequest, errBind.Error())
	}

	errVal := c.Validate(userRequest)
	if errVal != nil {
		return echo.NewHTTPError(http.StatusBadRequest, errVal.Error())
	}

	err := controller.UserService.Verify(userRequest)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	return c.JSON(http.StatusOK, ch.ResponseOkNoData("account is verified"))
}

func (controller *UserControllerImpl) RefreshVerCode(c echo.Context) error {
	var userRequest dto.UserVerCodeRequest
	errBind := c.Bind(&userRequest)
	if errBind != nil {
		return echo.NewHTTPError(http.StatusBadRequest, errBind.Error())
	}

	errVal := c.Validate(userRequest)
	if errVal != nil {
		return echo.NewHTTPError(http.StatusBadRequest, errVal.Error())
	}

	err := controller.UserService.RefreshVerCode(userRequest)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	return c.JSON(http.StatusOK, ch.ResponseOkNoData("new email verification code sent"))
}
