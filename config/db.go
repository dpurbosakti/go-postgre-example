package config

import (
	"go.uber.org/zap"
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

func InitDb(c *Config, logger *zap.Logger) (DB *gorm.DB) {

	if c.DbConf.Dsn == "" {
		logger.Info("instantiation",
			zap.String("type", "db"),
			zap.String("source", "gorm"),
			zap.String("status", "unset"),
			zap.String("message", "no DSN provided"))
		return
	}
	if c.DbConf.Dialect != "mysql" && c.DbConf.Dialect != "postgres" {
		logger.Info("instantiation",
			zap.String("type", "db"),
			zap.String("source", "gorm"),
			zap.String("status", "unset"),
			zap.String("message", "no proper dialect provided"))
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
		logger.Panic("instantiation", zap.String("type", "db"), zap.String("source", "gorm"), zap.String("status", "panic"))
		panic("Failed to connect to database!")
	} else {
		DB = db
		logger.Info("instantiation",
			zap.String("type", "db"),
			zap.String("source", "gorm"),
			zap.String("status", "done"),
			zap.String("name", "db"))
	}

	return DB
}
