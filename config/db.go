package config

import (
	logger "github.com/sirupsen/logrus"
	"gorm.io/driver/mysql"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type DbConf struct {
	Dsn          string
	MaxOpenConns int
	MaxIdleConns int
	MaxIdleTime  string
	Dialect      string
	DsnTest      string
}

func InitDb(c *Config) (DB *gorm.DB) {

	if c.DbConf.Dsn == "" {
		logger.WithFields(logger.Fields{
			"type":    "db",
			"source":  "gorm",
			"status":  "unset",
			"message": "no DSN provided",
		}).Error("instantiation")
		return
	}
	if c.DbConf.Dialect != "mysql" && c.DbConf.Dialect != "postgres" {
		logger.WithFields(logger.Fields{
			"type":    "db",
			"source":  "gorm",
			"status":  "unset",
			"message": "no proper dialect provided",
		}).Info("instantiation")
		return
	}

	var gormD gorm.Dialector
	if c.DbConf.Dialect == "mysql" {
		gormD = mysql.Open(c.DbConf.Dsn)
	} else if c.DbConf.Dialect == "postgres" {
		gormD = postgres.Open(c.DbConf.Dsn)
	}

	db, err := gorm.Open(gormD, &gorm.Config{})
	if err != nil {
		logger.WithFields(logger.Fields{
			"type":    "db",
			"source":  "gorm",
			"status":  "panic",
			"message": "Failed to connect to database!",
		}).Error("instantiation")
		logger.Panic(err)
	} else {
		DB = db
		logger.WithFields(logger.Fields{
			"type":   "db",
			"source": "gorm",
			"status": "done",
		}).Info("instantiation")
	}

	return DB
}

func InitDBTest() *gorm.DB {

	db, e := gorm.Open(mysql.Open("root:1amnohero@tcp(localhost:3306)/learnechotest?parseTime=true"), &gorm.Config{})
	if e != nil {
		panic(e)
	}

	return db
}
