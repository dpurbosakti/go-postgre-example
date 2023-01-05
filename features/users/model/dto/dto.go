package dto

import (
	"time"

	"gorm.io/gorm"
)

type UserCreateRequest struct {
	Name     string `json:"name" mod:"trim" validate:"required"`
	Email    string `json:"email" validate:"required"`
	Password string `json:"password" mod:"trim" validate:"required"`
	Phone    string `json:"phone" validate:"required"`
	Address  string `json:"address" validate:"required"`
	Role     string `json:"role" validate:"required"`
}

type UserCreateResponse struct {
	ID        uint
	Name      string
	Email     string
	Phone     string
	Address   string
	Role      string
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt
}
