package service

import (
	"gopkg.in/gomail.v2"
	"os"
	"strconv"
)

type EmailService interface {
	SendEmailWithText(toEmail string, message string) error
}

type emailService struct {
}

func NewEmailService() EmailService {
	return &emailService{}
}

func (service *emailService) SendEmailWithText(toEmail string, message string) error {
	go func() {
		mailer := gomail.NewMessage()
		mailer.SetHeader("From", os.Getenv("MAIL_FROM_ADDRESS"))
		mailer.SetHeader("To", toEmail)
		mailer.SetHeader("Subject", "Test Bang")
		mailer.SetBody("text/html", message)

		port, err := strconv.Atoi(os.Getenv("MAIL_PORT"))
		if err != nil {
			panic(err)
		}

		dialer := gomail.NewDialer(
			os.Getenv("MAIL_HOST"),
			port,
			os.Getenv("MAIL_USERNAME"),
			os.Getenv("MAIL_PASSWORD"),
		)

		err = dialer.DialAndSend(mailer)
		if err != nil {
			panic(err)
		}
	}()

	return nil
}
