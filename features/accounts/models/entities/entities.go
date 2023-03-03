package entities

import (
	"learn-echo/features/accounts/models/domain"
	"learn-echo/features/accounts/models/dto"
	"learn-echo/pkg/pagination"

	"github.com/labstack/echo/v4"

	"gorm.io/gorm"
)

type AccountRepository interface {
	Create(tx *gorm.DB, input domain.Account) (domain.Account, error)
	GetDetail(tx *gorm.DB, userId uint) (domain.Account, error)
	GetList(tx *gorm.DB, pagination pagination.Pagination) (pagination.Pagination, error)
	Delete(tx *gorm.DB, userId uint) error
}

type AccountService interface {
	Create(input dto.AccountCreateRequest, userId uint) (result dto.AccountResponse, err error)
	GetDetail(userId uint) (result dto.AccountResponse, err error)
	GetList(page pagination.Pagination) (result pagination.Pagination, err error)
	Delete(userId int) (err error)
}

type AccountHandler interface {
	Create(c echo.Context) error
	GetDetail(c echo.Context) error
	GetList(c echo.Context) error
}
