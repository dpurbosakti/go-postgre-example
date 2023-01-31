package domain

import (
	"time"

	"gorm.io/gorm"
)

type User struct {
	Id         uint           `json:"id" gorm:"primaryKey"`
	Nik        string         `json:"nik" gorm:"unique"`
	Name       string         `json:"name" gorm:"type:varchar(100)"`
	Email      string         `json:"email" gorm:"unique"`
	Password   string         `json:"password" gorm:"type:varchar(300)"`
	Phone      string         `json:"phone" gorm:"unique"`
	Address    string         `json:"address" gorm:"type:varchar(100)"`
	Role       string         `json:"role" gorm:"type:varchar(100)"`
	VerCode    string         `json:"verCode"` //verification code
	IsVerified bool           `json:"isVerified"`
	CreatedAt  time.Time      `json:"createdAt"`
	UpdatedAt  time.Time      `json:"updatedAt"`
	DeletedAt  gorm.DeletedAt `json:"deletedAt" gorm:"index"`
}
