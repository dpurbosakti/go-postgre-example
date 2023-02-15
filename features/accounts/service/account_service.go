package service

import (
	"learn-echo/features/accounts/model/dto"
)

type AccountService interface {
	Create(input dto.AccountCreateRequest, userId uint) (result dto.AccountResponse, err error)
}
