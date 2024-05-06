package models

import (
	"errors"
	"regexp"
	"time"

	"github.com/google/uuid"
)

type User struct {
	ID               uuid.UUID `gorm:"type:uuid;default:uuid_generate_v4();primary_key"`
	FirstName        string    `gorm:"type:varchar(255);not null"`
	LastName         string    `gorm:"type:varchar(255);not null"`
	Password         string    `gorm:"not null"`
	Provider         string    `gorm:"not null"`
	UserName         string    `gorm:"type:varchar(255);unique;not null"`
	Role             string    `gorm:"type:varchar(255);not null"`
	Email            string    `gorm:"uniqueIndex;unique;not null"`
	Photo            string    `gorm:"type:varchar(255);default null"`
	Verified         bool      `gorm:"not null"`
	VerificationCode string
	CreatedAt        time.Time `gorm:"not null"`
	UpdatedAt        time.Time `gorm:"not null"`
}

type SignUpInput struct {
	FirstName       string `json:"firstName" binding:"required"`
	LastName        string `json:"lastName" binding:"required"`
	Email           string `json:"email" binding:"required"`
	UserName        string `json:"username" binding:"required"`
	Password        string `json:"password" binding:"required,min=8"`
	PasswordConfirm string `json:"passwordConfirm" binding:"required"`
	Photo           string `json:"photo"`
}

type SignInInput struct {
	UserName string `json:"userName"`
	Email    string `json:"email"`
	Password string `json:"password" binding:"required"`
}

type UserResponse struct {
	ID        uuid.UUID `json:"id,omitempty"`
	FirstName string    `json:"firstName,omitempty"`
	LastName  string    `json:"lastName,omitempty"`
	UserName  string    `json:"userName", omitempty"`
	Email     string    `json:"email,omitempty"`
	Role      string    `json:"role,omitempty"`
	Photo     string    `json:"photo,omitempty"`
	Provider  string    `json:"provider"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func (input *SignInInput) Validate() error {
	if input.UserName == "" && input.Email == "" {
		return errors.New("either userName or email is required")
	}

	if input.Email != "" && !isValidEmail(input.Email) {
		return errors.New("email is not valid")
	}

	return nil
}

func isValidEmail(email string) bool {
	emailRegex := `^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`
	match, _ := regexp.MatchString(emailRegex, email)
	return match
}
