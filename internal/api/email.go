package api

import (
	"fmt"
	"net/smtp"
	"os"
)

func sendEmail(to string, rate float32) error {
	msg := []byte(fmt.Sprintf("To: %s\r\n"+
		"Subject: Daily USD to UAH exchange rate\r\n"+
		"\r\n"+
		"Current USD to UAH exchange rate is: %f \r\n",
		to, rate))
	smtpAddress := fmt.Sprintf("%s:%s", os.Getenv("SMTP_HOST"), os.Getenv("SMTP_PORT"))
	return smtp.SendMail(smtpAddress,
		nil, os.Getenv("SMTP_USERNAME"), []string{to}, msg)
}
