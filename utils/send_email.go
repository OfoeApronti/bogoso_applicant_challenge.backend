package utils

import (
	"math/rand"
	"net/smtp"
	"time"

	"fmt"

	"blow.com/bogoso.backend/app"
)

func SendEmail(recipient_email, token string) error {
	from := "ofoe.apronti@gmail.com"
	password := app.GmailAppToken
	to := []string{recipient_email}
	smtpHost := "smtp.gmail.com"
	smtpPort := "587"

	message := []byte(fmt.Sprintf("From: %s\r\n"+
		"To: %s\r\n"+
		"Subject: Validate email\r\n\r\n"+
		"Hi, \r\n, validate your email with the token: %s", from, to, token))

	// Create authentication
	auth := smtp.PlainAuth("", from, password, smtpHost)

	// Send actual message
	fmt.Println("sending the mail")
	err := smtp.SendMail(smtpHost+":"+smtpPort, auth, from, to, message)
	if err != nil {
		return err
	}
	return nil
}

func GetRandomNumber() string {
	rand.Seed(time.Now().UnixNano())
	randomInt := rand.Intn(9999)
	return fmt.Sprint(randomInt)
}
