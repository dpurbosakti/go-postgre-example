package routes

import (
	"learn-echo/factory"
	"net/http"

	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
)

type CustomValidator struct {
	validator *validator.Validate
}

func (cv *CustomValidator) Validate(i interface{}) error {
	if err := cv.validator.Struct(i); err != nil {
		// Optionally, you could return the error to give each route more control over the status code
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return nil
}

func New(presenter factory.Presenter) *echo.Echo {
	e := echo.New()

	e.Validator = &CustomValidator{validator: validator.New()}

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
