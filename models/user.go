package models

import (
	"time"

	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type User struct {
	ID                  uint       `json:"id" gorm:"primaryKey;autoIncrement"`
	Status              string     `json:"status" gorm:"not null"`
	Username            string     `json:"username" gorm:"unique;not null"`
	Email               string     `json:"email" gorm:"unique;not null"`
	Password            string     `json:"-" gorm:"not null"`
	Phone               string     `json:"phone" gorm:"not null"`
	Token               *string    `json:"token,omitempty" gorm:"default:null"`
	OtpVerify           int        `json:"otp_verify,omitempty" gorm:"default:null"`
	OtpReminder         *time.Time `json:"otp_reminder,omitempty" gorm:"default:null"`
	OtpPassword         int        `json:"otp_password,omitempty" gorm:"default:null"`
	OtpPasswordReminder *time.Time `json:"otp_password_reminder,omitempty" gorm:"default:null"`
	CreatedAt           time.Time  `json:"created_at"`
	UpdatedAt           time.Time  `json:"updated_at"`
}

func (User) TableName() string {
	return "users"
}

func (u *User) BeforeCreate(tx *gorm.DB) error {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(u.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	u.Password = string(hashedPassword)
	return nil
}

func (u *User) CheckPassword(password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(u.Password), []byte(password))
	return err == nil
}
