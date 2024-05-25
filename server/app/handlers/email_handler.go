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
func sendEmail(to, body, subject string) error {
	m := gomail.NewMessage()
	m.SetHeader("From", "no-reply@example.com")
	m.SetHeader("To", to)
	m.SetHeader("Subject", subject)
	m.SetBody("text/html", body)

	d := gomail.NewDialer("localhost", 1025, "", "")

	if err := d.DialAndSend(m); err != nil {
		return err
	}

	return nil
}
func (handler *EmailHandlerImpl) SendWelcomeEmail(payload *apis.WelcomeEmail) error {
	body := fmt.Sprintf("Welcome, %s", payload.Name)
	if err := sendEmail(payload.To, body, "Welcome"); err != nil {
		return err
	}

	return nil
}

func (handler *EmailHandlerImpl) SendForgotPasswordEmail(payload *apis.ForgotPasswordEmail) error {
	body := fmt.Sprintf("<h1>Welcome, This is your password reset link %s</h1>", payload.Code)
	if err := sendEmail(payload.Email, body, "Password Reset"); err != nil {
		return err
	}
	return nil
}
