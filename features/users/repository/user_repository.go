package repository

import (
	"learn-echo/features/users/models/domain"
	"learn-echo/features/users/models/dto"
	"learn-echo/pkg/pagination"

	"gorm.io/gorm"
)

type UserRepository interface {
	Create(tx *gorm.DB, input domain.User) (domain.User, error)
	Login(tx *gorm.DB, input dto.UserLoginRequest) (domain.User, error)
	GetDetail(tx *gorm.DB, userId int) (domain.User, error)
	GetList(tx *gorm.DB, pagination pagination.Pagination) (pagination.Pagination, error)
	Update(tx *gorm.DB, input domain.User) (domain.User, error)
	Delete(tx *gorm.DB, userId int) error
	CheckDuplicate(tx *gorm.DB, input domain.User) error
	CheckEmail(tx *gorm.DB, email string) (domain.User, error)
}
