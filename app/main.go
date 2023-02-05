package main

import (
	"fmt"
	"learn-echo/config"
	"learn-echo/factory"
	"learn-echo/middlewares"
	"learn-echo/migration"
	"learn-echo/routes"
	"os"
	"os/signal"
	"runtime/pprof"

	"github.com/labstack/echo/v4/middleware"
	log "github.com/sirupsen/logrus"
)

func main() {
	config.GetConfig()
	db := config.InitDb(config.Cfg)
	migration.InitMigrate(db)

	presenter := factory.InitFactory(db)
	e := routes.New(presenter)

	e.Use(middlewares.LogMiddleware)
	e.Use(middlewares.CorsMiddleware())
	e.Pre(middleware.RemoveTrailingSlash())

	//add log when c signal sent
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	go func() {
		for sig := range c {
			log.WithFields(log.Fields{
				"status": "closed",
				"signal": sig,
			}).Info("Program closed")
			pprof.StopCPUProfile()
			os.Exit(1)
		}
	}()

	e.Logger.Fatal(e.Start(fmt.Sprintf(":%d", config.Cfg.HttpConf.Port)))
}
