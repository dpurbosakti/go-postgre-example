package service

import (
	"fmt"
	"learn-echo/features/accounts/model/domain"
	"learn-echo/features/accounts/model/dto"
	"learn-echo/features/accounts/repository"

	"gorm.io/gorm"
)

type AccountServiceImpl struct {
	AccountRepository repository.AccountRepository
	DB                *gorm.DB
}

func NewAccountService(accountRepository repository.AccountRepository, db *gorm.DB) AccountService {
	return &AccountServiceImpl{
		AccountRepository: accountRepository,
		DB:                db,
	}
}

func (service *AccountServiceImpl) Create(input dto.AccountCreateRequest, userId uint) (result dto.AccountResponse, err error) {
	dataAccount := domain.Account{
		Type:    input.Type,
		Name:    fmt.Sprintf("bank account user id %d", userId),
		Balance: 0,
		User_Id: userId,
	}

	err = service.DB.Transaction(func(tx *gorm.DB) error {
		resultRepo, err := service.AccountRepository.Create(tx, dataAccount)
		if err != nil {
			return err
		}
		result = modelToResponse(resultRepo)
		return nil
	})
	if err != nil {
		return result, err
	}

	return result, nil
}