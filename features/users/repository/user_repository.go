package repository

import (
	"learn-echo/features/users/model/domain"
	"learn-echo/features/users/model/dto"

	"gorm.io/gorm"
)

type UserRepository interface {
	Create(tx *gorm.DB, input domain.User) (domain.User, error)
	Login(tx *gorm.DB, input dto.UserLoginRequest) (domain.User, error)
	// Update(tx *gorm.DB, input domain.User) (any, error)
	// GetList(tx *gorm.DB) (any, error)
	// GetDetail(tx *gorm.DB) (any, error)
	// Delete(tx *gorm.DB) (any, error)
}
