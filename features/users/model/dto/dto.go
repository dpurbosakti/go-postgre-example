package dto

import (
	"time"

	"gorm.io/gorm"
)

type UserCreateRequest struct {
	Name     string `json:"name" mod:"trim" validate:"required"`
	Nik      string `json:"nik" validate:"nik"`
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" mod:"trim" validate:"required"`
	Phone    string `json:"phone" validate:"required"`
	Address  string `json:"address" validate:"required"`
	Role     string `json:"role" validate:"required"`
}

type UserCreateResponse struct {
	ID        uint           `json:"id"`
	Name      string         `json:"name"`
	Nik       string         `json:"nik"`
	Email     string         `json:"email"`
	Phone     string         `json:"phone"`
	Address   string         `json:"address"`
	Role      string         `json:"role"`
	CreatedAt time.Time      `json:"createdAt"`
	UpdatedAt time.Time      `json:"updatedAt"`
	DeletedAt gorm.DeletedAt `json:"deletedAt"`
}
