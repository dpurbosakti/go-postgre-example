package migration

import (
	"learn-echo/features/users/model/domain"

	"gorm.io/gorm"
)

var autoMigrateList = []interface{}{
	&domain.User{},
}

func InitMigrate(Db *gorm.DB) {
	Db.AutoMigrate(autoMigrateList...)
}
