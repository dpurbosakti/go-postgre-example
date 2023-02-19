package repository

import (
	"learn-echo/features/accounts/model/domain"
	"learn-echo/pkg/pagination"

	"gorm.io/gorm"
)

type AccountRepository interface {
	Create(tx *gorm.DB, input domain.Account) (domain.Account, error)
	GetDetail(tx *gorm.DB, userId uint) (domain.Account, error)
	GetList(tx *gorm.DB, pagination pagination.Pagination) (pagination.Pagination, error)
}
