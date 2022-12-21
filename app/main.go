package main

import (
	"fmt"
	"learn-echo/config"
	"learn-echo/migration"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	cfg := config.GetConfig()
	logger := config.InitLogger()
	db := config.InitDb(cfg, logger)
	migration.InitMigrate(db)

	// mysql.DBMigration(db)

	e := echo.New()
	e.GET("/home", func(c echo.Context) error {
		data := map[string]interface{}{
			"message": "Welcome !!",
		}

		return c.JSON(http.StatusOK, data)
	})

	e.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
		Format: `[${time_rfc3339}] ${status} ${method} ${host}${path} ${latency_human}` + "\n",
	}))

	// factory.InitFactory(e, db)

	e.Logger.Fatal(e.Start(fmt.Sprintf(":%d", cfg.HttpConf.Port)))
}
