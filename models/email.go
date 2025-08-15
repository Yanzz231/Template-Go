// models/email.go
package models

import (
	"fmt"
	"os"

	"gopkg.in/gomail.v2"
)

type EmailService struct {
	SMTPHost     string
	SMTPPort     int
	SMTPUsername string
	SMTPPassword string
	FromEmail    string
	FromName     string
}

func (e *EmailService) SendOTP(email string, otp int, emailType string) error {
	isReset := emailType == "forget_password"

	subject := "Verifikasi email kamu yuk!"
	if isReset {
		subject = "Reset Password"
	}

	actionURL := fmt.Sprintf("%s/verify?email=%s&type=%s", os.Getenv("WEBSITE"), email, emailType)

	title := "Verifikasi Email Anda"
	message := "Klik tombol di bawah ini untuk memverifikasi email Anda."
	if isReset {
		title = "Reset Password"
		message = "Klik tombol di bawah ini untuk reset password Anda"
	}

	html := fmt.Sprintf(`
		<html>
			<body style="font-family: Arial, sans-serif; background-color: #f4f4f4; margin: 0; padding: 20px;">
				<div style="max-width: 600px; margin: 0 auto; background-color: #ffffff; padding: 20px; border-radius: 10px; box-shadow: 0px 4px 6px rgba(0, 0, 0, 0.1);">
					<h1 style="color: #333333; text-align: center;">%s</h1>
					<p style="color: #555555; text-align: center;">%s</p>
					<a href="%s" style="display: block; text-align: center; background-color: #ffcc00; color: #ffffff; font-size: 16px; padding: 10px 20px; border-radius: 5px; text-decoration: none;">Verifikasi Email</a>
					<p style="text-align: center; color: #555555;">Kode OTP Anda: <b>%s</b></p>
				</div>
			</body>
		</html>
	`, title, message, actionURL, otp)

	m := gomail.NewMessage()
	m.SetHeader("From", e.FromEmail)
	m.SetHeader("To", email)
	m.SetHeader("Subject", subject)
	m.SetBody("text/html", html)

	d := gomail.NewDialer(e.SMTPHost, e.SMTPPort, e.SMTPUsername, e.SMTPPassword)
	return d.DialAndSend(m)
}
