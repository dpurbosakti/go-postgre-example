package config

import (
	log "github.com/sirupsen/logrus"
	"gorm.io/driver/mysql"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
)

type DbConf struct {
	Dsn          string
	MaxOpenConns int
	MaxIdleConns int
	MaxIdleTime  string
	Dialect      string
}

func InitDb(c *Config) (DB *gorm.DB) {

	if c.DbConf.Dsn == "" {
		log.WithFields(log.Fields{
			"type":    "db",
			"source":  "gorm",
			"status":  "unset",
			"message": "no DSN provided",
		}).Info("Instantiation")
		return
	}
	if c.DbConf.Dialect != "mysql" && c.DbConf.Dialect != "postgres" {
		log.WithFields(log.Fields{
			"type":    "db",
			"source":  "gorm",
			"status":  "unset",
			"message": "no proper dialect provided",
		}).Info("Instantiation")
		return
	}

	var gormD gorm.Dialector
	if c.DbConf.Dialect == "mysql" {
		gormD = mysql.Open(c.DbConf.Dsn)
	} else if c.DbConf.Dialect == "postgres" {
		gormD = postgres.Open(c.DbConf.Dsn)
	}

	db, err := gorm.Open(gormD, &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			SingularTable: true,
			NoLowerCase:   true,
		},
	})
	if err != nil {
		log.WithFields(log.Fields{
			"type":    "db",
			"source":  "gorm",
			"status":  "panic",
			"message": "Failed to connect to database!",
		}).Info("Instantiation")
		log.Panic(err)
	} else {
		DB = db
		log.WithFields(log.Fields{
			"type":   "db",
			"source": "gorm",
			"status": "done",
		}).Info("Instantiation")
	}

	return DB
}
