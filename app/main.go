package main

import (
	"fmt"
	"learn-echo/config"
	"learn-echo/cron"
	"learn-echo/factory"
	"learn-echo/migration"
	"learn-echo/routes"
	"os"
	"os/signal"
	"runtime/pprof"

	"github.com/sirupsen/logrus"
)

func main() {
	config.GetConfig()

	db := config.InitDb(config.Cfg)
	migration.InitMigrate(db)

	presenter := factory.InitFactory(db)
	e := routes.New(presenter)

	cron.InitCron(db)

	//add log when c signal sent
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	go func() {
		for sig := range c {
			config.Cfg.LoggerConf.WithFields(logrus.Fields{
				"status": "closed",
				"signal": sig,
			}).Info("program closed")
			pprof.StopCPUProfile()
			os.Exit(1)
		}
	}()

	e.Logger.Fatal(e.Start(fmt.Sprintf(":%d", config.Cfg.HttpConf.Port)))
}
