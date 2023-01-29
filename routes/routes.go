package routes

import (
	"learn-echo/factory"
	"learn-echo/middlewares"
	"learn-echo/validation"
	"net/http"

	"github.com/labstack/echo/v4"
)

func New(presenter factory.Presenter) *echo.Echo {
	e := echo.New()

	e.Validator = &validation.CustomValidator{Validator: validation.InitValidator()}

	//index
	e.GET("/home", func(c echo.Context) error {
		data := map[string]interface{}{
			"message": "Welcome !!",
		}

		return c.JSON(http.StatusOK, data)
	})

	//users
	e.POST("/signup", presenter.UserPresenter.Create)
	e.POST("/login", presenter.UserPresenter.Login)
	e.GET("/user", presenter.UserPresenter.GetDetail, middlewares.IsAuthenticated())
	e.DELETE("/user", presenter.UserPresenter.Delete, middlewares.IsAuthenticated())
	e.PUT("/user", presenter.UserPresenter.Update, middlewares.IsAuthenticated())
	e.GET("/users", presenter.UserPresenter.GetList)

	return e
}
