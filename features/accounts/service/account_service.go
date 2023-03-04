package service

import (
	"fmt"
	"learn-echo/features/accounts/models/domain"
	"learn-echo/features/accounts/models/dto"
	"learn-echo/features/accounts/models/entities"
	"learn-echo/pkg/pagination"
	sh "learn-echo/pkg/servicehelper"
	"strings"

	"gorm.io/gorm"
)

const scope = "account"

type AccountService struct {
	AccountRepository entities.AccountRepository
	DB                *gorm.DB
}

func NewAccountService(accountRepository entities.AccountRepository, db *gorm.DB) entities.AccountService {
	return &AccountService{
		AccountRepository: accountRepository,
		DB:                db,
	}
}

func (as *AccountService) Create(input dto.AccountCreateRequest, userId uint) (result dto.AccountResponse, err error) {
	dataAccount := domain.Account{
		Type:    input.Type,
		Name:    fmt.Sprintf("bank account user id %d", userId),
		Balance: 0,
		User_Id: userId,
	}

	err = as.DB.Transaction(func(tx *gorm.DB) error {
		resultRepo, err := as.AccountRepository.Create(tx, dataAccount)
		if err != nil {
			return err
		}
		result = modelToResponse(resultRepo)
		return nil
	})
	if err != nil {
		err = sh.SetError(scope, "create", "error create new data", err.Error())
		return result, err
	}

	return result, nil
}

func (as *AccountService) GetDetail(userId uint) (result dto.AccountResponse, err error) {
	err = as.DB.Transaction(func(tx *gorm.DB) error {
		resultRepo, err := as.AccountRepository.GetDetail(tx, userId)
		if err != nil {
			return err
		}
		result = modelToResponse(resultRepo)
		return nil
	})
	if err != nil {
		err = sh.SetError(scope, "get-detail", "error get-detail", err.Error())
		return result, err
	}

	return result, nil
}

func (as *AccountService) GetList(page pagination.Pagination) (result pagination.Pagination, err error) {
	if page.Sort != "" {
		tmp := strings.Replace(page.Sort, "_", " ", 1)
		page.Sort = tmp
	}
	err = as.DB.Transaction(func(tx *gorm.DB) error {
		resultRepo, err := as.AccountRepository.GetList(tx, page)
		if err != nil {
			return err
		}
		result = resultRepo
		return nil
	})
	if err != nil {
		err = sh.SetError(scope, "get-list", "error get-list", err.Error())
		return result, err
	}

	return result, nil
}

func (as *AccountService) Delete(userId int) (err error) {
	err = as.DB.Transaction(func(tx *gorm.DB) error {
		err := as.AccountRepository.Delete(tx, uint(userId))
		if err != nil {
			return err
		}
		return nil
	})
	if err != nil {
		err = sh.SetError(scope, "delete", "error delete data", err.Error())
		return err
	}

	return nil
}
