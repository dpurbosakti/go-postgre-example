package emailhelper

import (
	"bytes"
	"fmt"
	"html/template"
	"learn-echo/config"
	"learn-echo/features/users/model/domain"

	"github.com/go-gomail/gomail"
)

func SendEmail(user domain.User) error {
	config := config.GetConfig()
	var body bytes.Buffer
	t, err := template.ParseFiles("../pkg/emailhelper/body.html")
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
