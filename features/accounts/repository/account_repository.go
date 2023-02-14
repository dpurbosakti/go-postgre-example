package repository

import (
	"learn-echo/features/accounts/model/domain"

	"gorm.io/gorm"
)

type AccountRepository interface {
	Create(tx *gorm.DB, input domain.Account) (domain.Account, error)
}
