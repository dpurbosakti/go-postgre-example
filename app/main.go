package main

import (
	"bytes"
	"fmt"
	"learn-echo/config"
	"learn-echo/factory"
	"learn-echo/features/users/model/domain"
	"learn-echo/middlewares"
	"learn-echo/migration"
	"learn-echo/routes"
	"os"
	"os/signal"
	"runtime/pprof"
	"text/template"

	"github.com/go-gomail/gomail"
	"github.com/robfig/cron"

	"github.com/labstack/echo/v4/middleware"
	log "github.com/sirupsen/logrus"
)

func main() {
	config.GetConfig()

	logger := log.New()
	logger.Formatter = &log.TextFormatter{
		ForceColors:     true,
		TimestampFormat: "2006/01/02 - 15:04:05",
		FullTimestamp:   true,
	}

	db := config.InitDb(config.Cfg, logger)
	migration.InitMigrate(db)

	presenter := factory.InitFactory(db)
	e := routes.New(presenter)

	e.Use(middlewares.LogMiddleware)
	e.Use(middlewares.CorsMiddleware())
	e.Pre(middleware.RemoveTrailingSlash())

	cron := cron.New()
	cron.AddFunc("1 * * * *", func() {
		var users []domain.User
		result := db.Find(&users)
		if result.Error != nil {
			fmt.Println("can't get user data")
		}

		config := config.Cfg
		var body bytes.Buffer
		t, err := template.ParseFiles("../pkg/emailhelper/info.html")
		if err != nil {
			fmt.Println(err.Error())
		}

		m := gomail.NewMessage()
		m.SetHeader("From", config.EmailConf.Email)
		m.SetHeader("Subject", "Information")
		t.Execute(&body, users)
		m.SetBody("text/html", body.String())

		for _, v := range users {
			m.SetHeader("To", v.Email)
			d := gomail.NewDialer(config.EmailConf.Host, config.EmailConf.Port, config.EmailConf.Email, config.EmailConf.Password)
			err = d.DialAndSend(m)
			if err != nil {
				fmt.Println("Error sending email:" + err.Error())

			}
		}
		fmt.Println("Email sent.")
	})
	cron.Start()

	//add log when c signal sent
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	go func() {
		for sig := range c {
			logger.WithFields(log.Fields{
				"status": "closed",
				"signal": sig,
			}).Info("Program closed")
			pprof.StopCPUProfile()
			os.Exit(1)
		}
	}()

	e.Logger.Fatal(e.Start(fmt.Sprintf(":%d", config.Cfg.HttpConf.Port)))
}
