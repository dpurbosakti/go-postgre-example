package middlewares

import (
	"context"
	"fmt"
	"learn-echo/database/redis"
	"learn-echo/features/users/models/dto"
	"strconv"
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
	duration := time.Hour * 168
	expired := time.Now().Add(duration).Unix()
	claims := jwt.MapClaims{}
	claims["authorized"] = true
	claims["userId"] = input.Id
	claims["role"] = input.Role
	claims["exp"] = expired
	claims["phone"] = input.Phone
	claims["email"] = input.Email
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// save to redis
	now := time.Now()
	rdExpired := time.Unix(expired, 0) //converting Unix to UTC(to Time object)
	errAccess := redis.RdClient.Set(context.Background(), strconv.Itoa(int(input.Id)), input.Id, rdExpired.Sub(now)).Err()
	if errAccess != nil {
		return "", errAccess
	}
	return token.SignedString([]byte("secret"))
}

func ExtractToken(e echo.Context) (result dto.UserDataToken, err error) {
	user := e.Get("user").(*jwt.Token)
	if user.Valid {
		claims := user.Claims.(jwt.MapClaims)
		result.Id = uint(claims["userId"].(float64))
		result.Role = claims["role"].(string)
		result.Phone = claims["phone"].(string)
		result.Email = claims["email"].(string)
		return result, nil
	}
	return result, fmt.Errorf("token invalid")
}
