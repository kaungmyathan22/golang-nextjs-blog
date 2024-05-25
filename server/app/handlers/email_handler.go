package handlers

import (
	"fmt"

	"github.com/kaungmyathan22/golang-nextjs-blog/app/models/apis"
	"gopkg.in/gomail.v2"
)

type EmailHandlerImpl struct{}

func NewEmailHandler() *EmailHandlerImpl {
	return &EmailHandlerImpl{}
}

func (handler *EmailHandlerImpl) SendWelcomeEmail(payload *apis.WelcomeEmail) error {
	m := gomail.NewMessage()
	m.SetHeader("From", "no-reply@example.com")
	m.SetHeader("To", payload.To)
	m.SetHeader("Subject", "Welcome")
	m.SetBody("text/html", fmt.Sprintf("<h1>Welcome, %s</h1>", payload.Name))

	d := gomail.NewDialer("localhost", 1025, "", "")

	if err := d.DialAndSend(m); err != nil {
		return err
	}

	return nil
}
