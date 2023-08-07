package cron

import (
	"bytes"
	"fmt"
	"learn-echo/config"
	"learn-echo/features/users/models/domain"
	"text/template"

	"github.com/go-gomail/gomail"
	"github.com/robfig/cron"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

func InitCron(db *gorm.DB) {
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
				// fmt.Println("Error sending email:" + err.Error())
				config.LoggerConf.Error(logrus.WithField("error sending email", err.Error()))

			}
		}
		config.LoggerConf.Info(logrus.WithField("email", "sent"))
	})
	cron.Start()
}
