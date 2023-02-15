package migration

import (
	account "learn-echo/features/accounts/model/domain"
	user "learn-echo/features/users/model/domain"

	"gorm.io/gorm"
)

func InitMigrate(Db *gorm.DB) {
	Db.AutoMigrate(&user.User{})
	Db.AutoMigrate(&account.Account{})
}
