// File: internal/domain/user/user.go

package user

import (
	"errors"
	"regexp"

	"golang.org/x/crypto/bcrypt"
)

type User struct {
	ID       uint   `json:"id"`
	Username string `json:"username"`
	Email    string `json:"email"`
	Password string `json:"-"`
}

func NewUser(username, email, password string) (*User, error) {
	if err := ValidateUsername(username); err != nil {
		return nil, err
	}
	if err := ValidateEmail(email); err != nil {
		return nil, err
	}
	if err := validatePassword(password); err != nil {
		return nil, err
	}

	hashedPassword, err := hashPassword(password)
	if err != nil {
		return nil, err
	}

	return &User{
		Username: username,
		Email:    email,
		Password: hashedPassword,
	}, nil
}

func (u *User) UpdatePassword(newPassword string) error {
	if err := validatePassword(newPassword); err != nil {
		return err
	}

	hashedPassword, err := hashPassword(newPassword)
	if err != nil {
		return err
	}

	u.Password = hashedPassword
	return nil
}

func ValidateUsername(username string) error {
	if len(username) < 3 || len(username) > 50 {
		return errors.New("username must be between 3 and 50 characters")
	}
	return nil
}

func ValidateEmail(email string) error {
	emailRegex := regexp.MustCompile(`^[a-z0-9._%+\-]+@[a-z0-9.\-]+\.[a-z]{2,4}$`)
	if !emailRegex.MatchString(email) {
		return errors.New("invalid email format")
	}
	return nil
}

func validatePassword(password string) error {
	if len(password) < 8 {
		return errors.New("password must be at least 8 characters long")
	}
	return nil
}

func hashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}

func CheckPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}
