package main

import (
	"bytes"
	"fmt"
	"learn-echo/config"
	"learn-echo/factory"
	"learn-echo/features/users/models/domain"
	"learn-echo/migration"
	"learn-echo/routes"
	"os"
	"os/signal"
	"runtime/pprof"
	"text/template"

	"github.com/go-gomail/gomail"
	"github.com/robfig/cron"

	"github.com/sirupsen/logrus"
)

func main() {
	config.GetConfig()

	db := config.InitDb(config.Cfg)
	migration.InitMigrate(db)

	presenter := factory.InitFactory(db)
	e := routes.New(presenter)

	cron := cron.New()
	cron.AddFunc("1 * * * *", func() {
		var users []domain.User
		result := db.Find(&users)
		if result.Error != nil {
			fmt.Println("can't get user data")
		}

		config := config.Cfg
		var body bytes.Buffer
		// t, err := template.ParseFiles("../pkg/emailhelper/info.html")
		t, err := template.ParseFiles("pkg/emailhelper/info.html")
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
		config.LoggerConf.Info(logrus.WithField("email", "sent"))
	})
	cron.Start()

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
