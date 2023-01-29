package migration

import (
	"learn-echo/features/users/model/domain"

	"gorm.io/gorm"
)

func InitMigrate(Db *gorm.DB) {
	Db.AutoMigrate(&domain.User{})
}
