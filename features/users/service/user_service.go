package service

import (
	"learn-echo/features/users/model/dto"
)

type UserService interface {
	Create(input dto.UserCreateRequest) (result dto.UserCreateResponse, err error)
	Login(input dto.UserLoginRequest) (result dto.UserDataToken, err error)
}
