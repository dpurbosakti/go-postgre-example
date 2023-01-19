package domain

import (
	"time"

	"gorm.io/gorm"
)

type User struct {
	Id        uint   `gorm:"primaryKey"`
	Nik       string `gorm:"unique"`
	Name      string `gorm:"type:varchar(100)"`
	Email     string `gorm:"unique"`
	Password  string `gorm:"type:varchar(300)"`
	Phone     string `gorm:"unique"`
	Address   string `gorm:"type:varchar(100)"`
	Role      string `gorm:"type:varchar(100)"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
}
