package repository

import (
	"fmt"
	"learn-echo/features/accounts/model/domain"

	"gorm.io/gorm"
)

const source = "account"

type AccountRepositoryImpl struct {
}

func NewAccountRepository() AccountRepository {
	return &AccountRepositoryImpl{}
}

func (repository AccountRepositoryImpl) Create(tx *gorm.DB, input domain.Account) (domain.Account, error) {
	result := tx.Create(&input)
	if result.Error != nil {
		return domain.Account{}, fmt.Errorf("failed to create data %s", source)
	}

	return input, nil
}

func (repository AccountRepositoryImpl) GetDetail(tx *gorm.DB, userId uint) (domain.Account, error) {
	var account *domain.Account
	result := tx.Where("user_id = ?", userId).First(&account)
	if result.Error != nil {
		return domain.Account{}, fmt.Errorf("account with user id %d not found", userId)
	}

	return *account, nil
}
