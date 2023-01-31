package emailhelper

import (
	"net/smtp"
)

func SendMailSimple(subject, body string, to []string, errCh chan error) {
	auth := smtp.PlainAuth(
		"",
		"altacapstonegroup3@gmail.com",
		"ppvxpgjmyirnqmdm",
		"smtp.gmail.com",
	)

	msg := "Subject: " + subject + "\n" + body

	err := smtp.SendMail(
		"smtp.gmail.com:587",
		auth,
		"altacapstonegroup3@gmail.com",
		to,
		[]byte(msg),
	)

	if err != nil {
		errCh <- err
		return
	}
	errCh <- nil
}
