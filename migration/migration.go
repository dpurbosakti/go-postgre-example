package migration

import (
	account "learn-echo/features/accounts/models/domain"
	user "learn-echo/features/users/models/domain"

	"gorm.io/gorm"
)

func InitMigrate(db *gorm.DB) {
	db.AutoMigrate(&user.User{})
	db.AutoMigrate(&account.Account{})
}
