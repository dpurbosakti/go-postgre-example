package controller

import (
	"learn-echo/features/accounts/model/dto"
	"learn-echo/features/accounts/service"
	"learn-echo/middlewares"
	ch "learn-echo/pkg/controllerhelper"
	"learn-echo/pkg/pagination"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

type AccountControllerImpl struct {
	AccountService service.AccountService
}

func NewAccountController(accountService service.AccountService) *AccountControllerImpl {
	return &AccountControllerImpl{
		AccountService: accountService,
	}
}

func (controller *AccountControllerImpl) Create(c echo.Context) error {
	var accountReq dto.AccountCreateRequest
	dataToken, _ := middlewares.ExtractToken(c)
	userId := uint(dataToken.Id)

	errBind := c.Bind(&accountReq)
	if errBind != nil {
		return echo.NewHTTPError(http.StatusBadRequest, errBind.Error())
	}

	errVal := c.Validate(accountReq)
	if errVal != nil {
		return echo.NewHTTPError(http.StatusBadRequest, errVal.Error())
	}

	result, err := controller.AccountService.Create(accountReq, userId)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	return c.JSON(http.StatusCreated, ch.ResponseOkWithData("create data account success", result))

}

func (controller *AccountControllerImpl) GetDetail(c echo.Context) error {
	dataToken, _ := middlewares.ExtractToken(c)
	userId := uint(dataToken.Id)

	result, err := controller.AccountService.GetDetail(userId)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	return c.JSON(http.StatusCreated, ch.ResponseOkWithData("get data account success", result))
}

func (controller *AccountControllerImpl) GetList(c echo.Context) error {
	var page pagination.Pagination
	limitInt, _ := strconv.Atoi(c.QueryParam("limit"))
	pageInt, _ := strconv.Atoi(c.QueryParam("page"))
	page.Limit = limitInt
	page.Page = pageInt
	page.Sort = c.QueryParam("sort")

	result, err := controller.AccountService.GetList(page)
	if err != nil {
		return echo.NewHTTPError(http.StatusNotFound, err)
	}

	if result.TotalRows == 0 {
		return echo.NewHTTPError(http.StatusNotFound, "data not found")
	}

	return c.JSON(http.StatusOK, ch.ResponseOkWithData("get data accounts success", result))
}
