package migration

import (
	account "learn-echo/features/accounts/models/domain"
	user "learn-echo/features/users/models/domain"

	"gorm.io/gorm"
)

func InitMigrate(Db *gorm.DB) {
	Db.AutoMigrate(&user.User{})
	Db.AutoMigrate(&account.Account{})
}
