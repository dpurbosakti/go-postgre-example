package domain

import (
	m "learn-echo/features/users/models/domain"
	"time"

	"gorm.io/gorm"
)

type Account struct {
	Id        uint           `json:"id" gorm:"primaryKey"`
	Type      string         `json:"type"`
	Name      string         `json:"name"`
	Balance   uint           `json:"balance"`
	User_Id   uint           `json:"user_id"`
	User      *m.User        `json:"user,omitempty" gorm:"foreignKey:User_Id"`
	CreatedAt time.Time      `json:"createdAt"`
	UpdatedAt time.Time      `json:"updatedAt"`
	DeletedAt gorm.DeletedAt `json:"deletedAt" gorm:"index"`
}
