package main

import (
	"fmt"
	"learn-echo/config"
	"learn-echo/database/fakers"

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
	fmt.Println(config.Cfg)
	db := config.InitDb(config.Cfg)
	err := DbSeed(db)
	if err != nil {
		fmt.Println(err.Error())
	}
	fmt.Println("\nseeding done")
}
