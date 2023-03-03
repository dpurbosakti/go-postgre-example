package repository

import (
	"fmt"
	"learn-echo/features/accounts/models/domain"
	"learn-echo/features/accounts/models/entities"
	"learn-echo/pkg/pagination"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

const source = "account"

type AccountRepository struct {
}

func NewAccountRepository() entities.AccountRepository {
	return &AccountRepository{}
}

func (ar AccountRepository) Create(tx *gorm.DB, input domain.Account) (domain.Account, error) {
	result := tx.Create(&input)
	if result.Error != nil {
		return domain.Account{}, fmt.Errorf("failed to create data %s", source)
	}

	return input, nil
}

func (ar AccountRepository) GetDetail(tx *gorm.DB, userId uint) (domain.Account, error) {
	var account *domain.Account
	result := tx.Where("user_id = ?", userId).First(&account)
	if result.Error != nil {
		return domain.Account{}, fmt.Errorf("account with user id %d not found", userId)
	}

	return *account, nil
}

func (ar AccountRepository) GetList(tx *gorm.DB, page pagination.Pagination) (pagination.Pagination, error) {
	var accounts []domain.Account

	tx.Preload(clause.Associations, func(tx *gorm.DB) *gorm.DB {
		return tx.Omit("Password")
	}).Scopes(pagination.Paginate(accounts, &page, tx)).Find(&accounts)
	page.Rows = accounts

	return page, nil
}

func (ar AccountRepository) Delete(tx *gorm.DB, userId uint) error {
	var account domain.Account
	result := tx.Where("user_id = ?", userId).Delete(&account)
	if result.Error != nil {
		return fmt.Errorf("user id %d not found", userId)
	}

	return nil
}
