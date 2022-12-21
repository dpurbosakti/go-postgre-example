package migration

import "gorm.io/gorm"

var autoMigrateList []interface{}

func InitMigrate(Db *gorm.DB) {
	Db.AutoMigrate(autoMigrateList...)
}
