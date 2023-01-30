package validation

import (
	"fmt"
	"net/http"
	"regexp"

	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
)

type CustomValidator struct {
	Validator *validator.Validate
}

func (cv *CustomValidator) Validate(i interface{}) error {
	if err := cv.Validator.Struct(i); err != nil {
		// Optionally, you could return the error to give each route more control over the status code
		report := []string{}
		errs := err.(validator.ValidationErrors)
		for _, err := range errs {
			switch err.Tag() {
			case "required":
				report = append(report, fmt.Sprintf("%s is required",
					err.Field()))
			case "email":
				report = append(report, fmt.Sprintf("%s is not valid email",
					err.Value()))
			case "gte":
				report = append(report, fmt.Sprintf("%s value must be greater than %s",
					err.Field(), err.Param()))
			case "lte":
				report = append(report, fmt.Sprintf("%s value must be lower than %s",
					err.Field(), err.Param()))
			case "nik":
				report = append(report, fmt.Sprintf("%s was not in correct format e.g 3514142206950001",
					err.Value()))
			case "phone":
				report = append(report, fmt.Sprintf("%s was not in correct format e.g 085865968948",
					err.Value()))
			}

		}
		return echo.NewHTTPError(http.StatusBadRequest, report)
	}
	return nil
}

func nikValidator(fl validator.FieldLevel) bool {
	nik := fl.Field().String()
	regex := regexp.MustCompile(`^(1[1-9]|21|[37][1-6]|5[1-3]|6[1-5]|[89][12])\d{2}\d{2}([04][1-9]|[1256][0-9]|[37][01])(0[1-9]|1[0-2])\d{2}\d{4}$`)

	result := regex.MatchString(nik)

	return result
}

func phoneValidator(fl validator.FieldLevel) bool {
	phone := fl.Field().String()
	regex := regexp.MustCompile(`^(\+62|62|0)[2-9][1-9][0-9]{6,10}$`)

	result := regex.MatchString(phone)

	return result
}

// func sortPagination(fl validator.FieldLevel) bool {
// 	sort := fl.Field().String()
// 	return !strings.Contains(sort, ";")
// }

func InitValidator() *validator.Validate {
	validate := validator.New()

	validate.RegisterValidation("nik", nikValidator)
	validate.RegisterValidation("phone", phoneValidator)
	// validate.RegisterValidation("sort", sortPagination)
	return validate
}

// var ValidationErrorHandler = func(err error, c echo.Context) {
// 	report, ok := err.(*echo.HTTPError)
// 	if !ok {
// 		report = echo.NewHTTPError(http.StatusInternalServerError, err.Error())
// 	}

// 	if castedObject, ok := err.(validator.ValidationErrors); ok {
// 		for _, err := range castedObject {
// 			switch err.Tag() {
// 			case "required":
// 				report.Message = fmt.Sprintf("%s is required",
// 					err.Field())
// 			case "email":
// 				report.Message = fmt.Sprintf("%s is not valid email",
// 					err.Field())
// 			case "gte":
// 				report.Message = fmt.Sprintf("%s value must be greater than %s",
// 					err.Field(), err.Param())
// 			case "lte":
// 				report.Message = fmt.Sprintf("%s value must be lower than %s",
// 					err.Field(), err.Param())
// 			case "nik":
// 				report.Message = fmt.Sprintf("%s value must 16 digit nik",
// 					err.Field())
// 			}

// 			break
// 		}
// 	}

// 	c.Logger().Error(report)
// 	c.JSON(report.Code, report)
// }
