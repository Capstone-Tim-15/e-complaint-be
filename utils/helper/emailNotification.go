package helper

import (
	"fmt"
	"os"
	"strconv"

	"github.com/go-gomail/gomail"
)

type Mail struct {
	Host     string
	Port     string
	Username string
	Password string
}

func SendEmailNotification(email string) error {
	mail := Mail{
		Host:     os.Getenv("SMTP_SERVER"),
		Port:     os.Getenv("SMTP_PORT"),
		Username: os.Getenv("SMTP_USERNAME"),
		Password: os.Getenv("SMTP_PASSWORD"),
	}

	body := fmt.Sprintf("Hello %s, this is a test email from E-Complaint", email)

	to := []string{email}

	m := gomail.NewMessage()
	m.SetHeader("From", mail.Username)
	m.SetHeader("To", to...)
	m.SetHeader("Subject", "GovComplaint Notification")
	m.SetBody("text/html", body)

	Port, _ := strconv.Atoi(mail.Port)

	dialer := gomail.NewDialer(
		mail.Host,
		Port,
		mail.Username,
		mail.Password,
	)

	if err := dialer.DialAndSend(m); err != nil {
		return err
	}
	return nil
}

func SendOTP(email string, otp string) error {
	mail := Mail{
		Host:     os.Getenv("SMTP_SERVER"),
		Port:     os.Getenv("SMTP_PORT"),
		Username: os.Getenv("SMTP_USERNAME"),
		Password: os.Getenv("SMTP_PASSWORD"),
	}

	body := fmt.Sprintf("Hello,<br><br>Thank you for using our service. Your One-Time Password (OTP) is: <strong>%s</strong>.<br><br>If you did not request this OTP, please ignore this email.<br><br>Best regards,<br>Your Service Team", otp)

	to := []string{email}

	m := gomail.NewMessage()
	m.SetHeader("From", mail.Username)
	m.SetHeader("To", to...)
	m.SetHeader("Subject", "GovComplaint Notification")
	m.SetBody("text/html", body)

	Port, _ := strconv.Atoi(mail.Port)

	dialer := gomail.NewDialer(
		mail.Host,
		Port,
		mail.Username,
		mail.Password,
	)

	if err := dialer.DialAndSend(m); err != nil {
		return err
	}
	return nil
}
