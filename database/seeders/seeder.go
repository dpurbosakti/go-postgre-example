package main

import (
	"learn-echo/config"
	"learn-echo/database/fakers"

	"github.com/sirupsen/logrus"

	"gorm.io/gorm"
)

type Seeder struct {
	Seeder interface{}
}

func RegisterSeeders(db *gorm.DB) []Seeder {
	return []Seeder{
		{Seeder: fakers.UserFaker(db)},
	}
}

func DbSeed(db *gorm.DB) error {
	for _, seeder := range RegisterSeeders(db) {
		err := db.Debug().Create(seeder.Seeder).Error
		if err != nil {
			return err
		}
	}
	return nil
}

func main() {
	config.GetConfig()
	logger := config.Cfg.LoggerConf.WithField("func", "seeder")
	// fmt.Println(config.Cfg)
	db := config.InitDb(config.Cfg)
	err := DbSeed(db)
	if err != nil {
		logger.WithError(err).Error("error when seeding")
	}
	logger.Info(logrus.WithField("seeding", "done"))
}
