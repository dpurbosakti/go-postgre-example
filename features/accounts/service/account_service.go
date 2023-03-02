package service

import (
	"learn-echo/features/accounts/model/dto"
	"learn-echo/pkg/pagination"
)

type AccountService interface {
	Create(input dto.AccountCreateRequest, userId uint) (result dto.AccountResponse, err error)
	GetDetail(userId uint) (result dto.AccountResponse, err error)
	GetList(page pagination.Pagination) (result pagination.Pagination, err error)
	Delete(userId int) (err error)
}
