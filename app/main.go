package main

import (
	"fmt"
	"learn-echo/config"
	"learn-echo/factory"
	"learn-echo/migration"
	"learn-echo/routes"

	"github.com/labstack/echo/v4/middleware"
)

func main() {
	cfg := config.GetConfig()
	logger := config.InitLogger()
	db := config.InitDb(cfg, logger)
	migration.InitMigrate(db)

	// mysql.DBMigration(db)

	// validate := validator.New()
	presenter := factory.InitFactory(db)
	e := routes.New(presenter)

	e.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
		Format: `[${time_rfc3339}] ${status} ${method} ${host}${path} ${latency_human}` + "\n",
	}))

	// factory.InitFactory(e, db)

	e.Logger.Fatal(e.Start(fmt.Sprintf(":%d", cfg.HttpConf.Port)))
}
