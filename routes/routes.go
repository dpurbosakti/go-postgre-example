package routes

import (
	"learn-echo/factory"
	"learn-echo/middlewares"
	"learn-echo/validation"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func New(presenter factory.Presenter) *echo.Echo {
	e := echo.New()

	e.Validator = &validation.CustomValidator{Validator: validation.InitValidator()}

	e.Use(middlewares.LogMiddleware)
	e.Use(middlewares.CorsMiddleware())
	e.Pre(middleware.RemoveTrailingSlash())

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
	e.GET("/users/detail", presenter.UserPresenter.GetDetail, middlewares.IsAuthenticated())
	e.DELETE("/users", presenter.UserPresenter.Delete, middlewares.IsAuthenticated())
	e.PUT("/users", presenter.UserPresenter.Update, middlewares.IsAuthenticated())
	e.GET("/users/list", presenter.UserPresenter.GetList)
	e.POST("/users/verify", presenter.UserPresenter.Verify)
	e.POST("/users/refreshcode", presenter.UserPresenter.RefreshVerCode)

	// accounts
	e.POST("/accounts", presenter.AccountPresenter.Create, middlewares.IsAuthenticated())
	e.GET("/accounts/detail", presenter.AccountPresenter.GetDetail, middlewares.IsAuthenticated())
	e.GET("/accounts/list", presenter.AccountPresenter.GetList)
	return e
}
