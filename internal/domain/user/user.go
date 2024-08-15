package user

import (
	"errors"
	"regexp"
	"strings"
)

// User represents the user entity in the domain
type User struct {
	Name     string
	Email    string
	Password string
}

// NewUser creates a new User instance
func NewUser(name, email, password string) (*User, error) {
	u := &User{
		Name:     name,
		Email:    email,
		Password: password,
	}

	if err := u.Validate(); err != nil {
		return nil, err
	}

	return u, nil
}

// Validate checks if the User fields are valid
func (u *User) Validate() error {
	if err := validateName(u.Name); err != nil {
		return err
	}

	if err := validateEmail(u.Email); err != nil {
		return err
	}

	if err := validatePassword(u.Password); err != nil {
		return err
	}

	return nil
}

func validateName(name string) error {
	if strings.TrimSpace(name) == "" {
		return errors.New("name cannot be empty")
	}
	return nil
}

func validateEmail(email string) error {
	if strings.TrimSpace(email) == "" {
		return errors.New("email cannot be empty")
	}

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
