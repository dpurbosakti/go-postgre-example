package service

import (
	"learn-echo/features/users/model/dto"
	"learn-echo/pkg/pagination"
)

type UserService interface {
	Create(input dto.UserCreateRequest) (result dto.UserResponse, err error)
	Login(input dto.UserLoginRequest) (result dto.UserDataToken, err error)
	GetDetail(userId int) (result dto.UserResponse, err error)
	GetList(page pagination.Pagination) (result pagination.Pagination, err error)
	Delete(userId int) (err error)
	Update(input dto.UserUpdateRequest, userId int) (result dto.UserResponse, err error)
	Verify(input dto.UserVerifyRequest) (err error)
}
