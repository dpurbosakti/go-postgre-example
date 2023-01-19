package routes

import (
	"learn-echo/factory"
	"learn-echo/validation"
	"net/http"

	"github.com/labstack/echo/v4"
)

func New(presenter factory.Presenter) *echo.Echo {
	e := echo.New()

	e.Validator = &validation.CustomValidator{Validator: validation.InitValidator()}

	//users
	e.POST("/user", presenter.UserPresenter.Create)
	e.GET("/home", func(c echo.Context) error {
		data := map[string]interface{}{
			"message": "Welcome !!",
		}

		return c.JSON(http.StatusOK, data)
	})
	return e
}
