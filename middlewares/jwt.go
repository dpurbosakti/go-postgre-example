package middlewares

import (
	"fmt"
	"learn-echo/features/users/model/dto"
	"time"

	"github.com/golang-jwt/jwt"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func IsAuthenticated() echo.MiddlewareFunc {
	return middleware.JWTWithConfig(middleware.JWTConfig{
		SigningMethod: "HS256",
		SigningKey:    []byte("secret"),
	})
}

func CreateToken(input dto.UserResponse) (string, error) {
	claims := jwt.MapClaims{}
	claims["authorized"] = true
	claims["userId"] = input.Id
	claims["role"] = input.Role
	claims["exp"] = time.Now().Add(time.Hour * 504).Unix()
	claims["handphone"] = input.Phone
	claims["email"] = input.Email
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte("secret"))
}

func ExtractToken(e echo.Context) (result dto.UserDataToken, err error) {
	user := e.Get("user").(*jwt.Token)
	if user.Valid {
		claims := user.Claims.(jwt.MapClaims)
		result.Id = claims["userId"].(uint)
		result.Role = claims["role"].(string)
		result.Phone = claims["phone"].(string)
		result.Email = claims["email"].(string)
		return result, nil
	}
	return result, fmt.Errorf("token invalid")
}
