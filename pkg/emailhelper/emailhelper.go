package emailhelper

import (
	"bytes"
	"fmt"
	"html/template"
	"learn-echo/config"
	"learn-echo/features/users/models/domain"
	"os"

	"github.com/go-gomail/gomail"
	"gorm.io/gorm"
)

func SendEmailVerCode(user domain.User) error {
	var basePath string
	wd, _ := os.Getwd()
	if string(wd[len(wd)-13]) == "u" {
		// basePath = filepath.Join(wd, "../../../", "pkg/emailhelper/body.html")
		basePath = "pkg/emailhelper/body.html"
	} else {
		// basePath = filepath.Join(wd, "../", "pkg/emailhelper/body.html")
		basePath = "pkg/emailhelper/body.html"
	}
	config := config.Cfg
	var body bytes.Buffer
	t, err := template.ParseFiles(basePath)
	if err != nil {
		return err
	}
	t.Execute(&body, user)
	m := gomail.NewMessage()
	m.SetHeader("From", config.EmailConf.Email)
	m.SetHeader("To", user.Email)
	m.SetHeader("Subject", "Verification Code")
	m.SetBody("text/html", body.String())
	d := gomail.NewDialer(config.EmailConf.Host, config.EmailConf.Port, config.EmailConf.Email, config.EmailConf.Password)
	err = d.DialAndSend(m)
	if err != nil {
		fmt.Println("Error sending email:" + err.Error())
		return err
	}
	fmt.Println("Email sent.")
	return nil
}

func SendEmail(db *gorm.DB) {
	var user []domain.User
	result := db.Find(&user)
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
	m.SetBody("text/html", body.String())
	d := gomail.NewDialer(config.EmailConf.Host, config.EmailConf.Port, config.EmailConf.Email, config.EmailConf.Password)
	for _, v := range user {
		m.SetHeader("To", v.Email)
		t.Execute(&body, v)
		err := d.DialAndSend(m)
		if err != nil {
			fmt.Println("Error sending email:" + err.Error())

		}
	}
	fmt.Println("Email sent.")
}
