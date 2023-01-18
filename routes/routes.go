package routes

import (
	"learn-echo/factory"
	"learn-echo/validation"
	"net/http"

	"github.com/labstack/echo/v4"
)

// type CustomValidator struct {
// 	validator *validator.Validate
// }

// func (cv *CustomValidator) Validate(i interface{}) error {
// 	if err := cv.validator.Struct(i); err != nil {
// 		// Optionally, you could return the error to give each route more control over the status code
// 		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
// 	}
// 	return nil
// }

func New(presenter factory.Presenter) *echo.Echo {
	e := echo.New()

	e.Validator = &validation.CustomValidator{Validator: validation.InitValidator()}
	// e.HTTPErrorHandler = validation.ValidationErrorHandler

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
