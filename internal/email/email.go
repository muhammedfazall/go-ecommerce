package email

import (
	"fmt"
	"os"
	"strconv"

	"gopkg.in/gomail.v2"
)

func SendEmail(to, subject, htmlBody string) error {
	port, _ := strconv.Atoi(os.Getenv("SMTP_PORT"))
	m := gomail.NewMessage()
	m.SetHeader("From", os.Getenv("SMTP_FROM"))
	m.SetHeader("To", to)
	m.SetHeader("Subject", subject)
	m.SetBody("text/html", htmlBody)
	d := gomail.NewDialer(
		os.Getenv("SMTP_HOST"),
		port,
		os.Getenv("SMTP_USER"),
		os.Getenv("SMTP_PASSWORD"),
	)
	return d.DialAndSend(m)
}

func SendWelcomeEmail(to, username string) error {
	subject := "Welcome to SneaCave!"
	body := fmt.Sprintf(`<h2>Hey %s, welcome to SneaCave!</h2>
        <p>Thanks for signing up. Start exploring the latest sneakers now.</p>`, username)
	return SendEmail(to, subject, body)
}

func SendOTPEmail(to, otp string) error {
    subject := "Your SneaCave Verification Code"
    body := fmt.Sprintf(`<h2>Email Verification</h2>
        <p>Your OTP code is: <strong>%s</strong></p>
        <p>This code expires in 5 minutes.</p>`, otp)
    return SendEmail(to, subject, body)
}