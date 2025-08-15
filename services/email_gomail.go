package services

import (
	"Template-Go/models"
	"os"
	"strconv"
)

func EmailServiceGomail() *models.EmailService {
	port, _ := strconv.Atoi(os.Getenv("SMTP_PORT"))

	return &models.EmailService{
		SMTPHost:     os.Getenv("SMTP_HOST"),
		SMTPPort:     port,
		SMTPUsername: os.Getenv("SMTP_USERNAME"),
		SMTPPassword: os.Getenv("SMTP_PASSWORD"),
		FromEmail:    os.Getenv("FROM_EMAIL"),
		FromName:     os.Getenv("FROM_NAME"),
	}
}
