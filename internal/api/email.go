package api

import (
	"fmt"
	"net/smtp"
	"os"
)

func sendEmail(to []string, msg []byte) error {
	auth := smtp.PlainAuth("", os.Getenv("SMTP_USERNAME"),
		os.Getenv("SMTP_PASSWORD"), os.Getenv("HOST"))

	smtpAddress := fmt.Sprintf("%s:%s", os.Getenv("SMTP_HOST"), os.Getenv("SMTP_PORT"))
	return smtp.SendMail(smtpAddress,
		auth, os.Getenv("SMTP_USERNAME"), to, msg)
}
