package service

import (
	"fmt"
	"learn-echo/features/accounts/model/domain"
	"learn-echo/features/accounts/model/dto"
	"learn-echo/features/accounts/repository"
	"learn-echo/pkg/pagination"
	sh "learn-echo/pkg/servicehelper"
	"strings"

	"gorm.io/gorm"
)

const scope = "account"

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

func (s *AccountServiceImpl) Create(input dto.AccountCreateRequest, userId uint) (result dto.AccountResponse, err error) {
	dataAccount := domain.Account{
		Type:    input.Type,
		Name:    fmt.Sprintf("bank account user id %d", userId),
		Balance: 0,
		User_Id: userId,
	}

	err = s.DB.Transaction(func(tx *gorm.DB) error {
		resultRepo, err := s.AccountRepository.Create(tx, dataAccount)
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

func (s *AccountServiceImpl) GetDetail(userId uint) (result dto.AccountResponse, err error) {
	err = s.DB.Transaction(func(tx *gorm.DB) error {
		resultRepo, err := s.AccountRepository.GetDetail(tx, userId)
		if err != nil {
			return err
		}
		result = modelToResponse(resultRepo)
		return nil
	})
	if err != nil {
		err = sh.SetError(scope, "create", "error get-detail", err.Error())
		return result, err
	}

	return result, nil
}

func (s *AccountServiceImpl) GetList(page pagination.Pagination) (result pagination.Pagination, err error) {
	if page.Sort != "" {
		tmp := strings.Replace(page.Sort, "_", " ", 1)
		page.Sort = tmp
	}
	err = s.DB.Transaction(func(tx *gorm.DB) error {
		resultRepo, err := s.AccountRepository.GetList(tx, page)
		if err != nil {
			return err
		}
		result = resultRepo
		return nil
	})
	if err != nil {
		err = sh.SetError(scope, "create", "error get-list", err.Error())
		return result, err
	}

	return result, nil
}
