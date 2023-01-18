package domain

import (
	"time"

	"gorm.io/gorm"
)

type User struct {
	ID        uint   `gorm:"primaryKey"`
	Nik       string `gorm:"type:varchar(16)"`
	Name      string `gorm:"type:varchar(100)"`
	Email     string `gorm:"type:varchar(100)"`
	Password  string `gorm:"type:varchar(100)"`
	Phone     string `gorm:"type:varchar(100)"`
	Address   string `gorm:"type:varchar(100)"`
	Role      string `gorm:"type:varchar(100)"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
}
