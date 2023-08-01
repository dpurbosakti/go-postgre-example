package dto

import (
	m "learn-echo/features/users/models/domain"
	"time"

	"gorm.io/gorm"
)

type AccountCreateRequest struct {
	Type string `json:"type" validate:"required"`
}

type AccountResponse struct {
	Id        uint           `json:"id"`
	Type      string         `json:"type"`
	Name      string         `json:"name"`
	Balance   uint           `json:"balance"`
	User_Id   uint           `json:"user_id"`
	User      *m.User        `json:"user,omitempty"`
	CreatedAt time.Time      `json:"createdAt"`
	UpdatedAt time.Time      `json:"updatedAt"`
	DeletedAt gorm.DeletedAt `json:"deletedAt" gorm:"index"`
}
