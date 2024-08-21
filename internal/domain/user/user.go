// File: internal/domain/user/user.go

package user

import (
	"errors"
	"regexp"
	"time"

	"golang.org/x/crypto/bcrypt"
)

type User struct {
	ID              uint      `gorm:"primaryKey" json:"id"`
	Username        string    `gorm:"not null;unique" json:"username"`
	FirstName       string    `gorm:"not null" json:"first_name"`
	LastName        string    `gorm:"not null" json:"last_name"`
	Email           string    `gorm:"not null;unique" json:"email"`
	Password        string    `gorm:"not null" json:"-"`
	IsActive        bool      `gorm:"not null;default:false" json:"is_active"`
	Role            string    `json:"role"`
	ProfileImageURL string    `json:"profile_image_url"`
	PhoneNumber     string    `json:"phone_number"`
	EmailVerified   bool      `gorm:"not null;default:false" json:"email_verified"`
	LastLogin       time.Time `json:"last_login"`
	CreatedAt       time.Time `gorm:"not null" json:"created_at"`
	UpdatedAt       time.Time `gorm:"not null" json:"updated_at"`
}

func NewUser(username, firstName, lastName, email, password, phoneNumber string) (*User, error) {
	if err := ValidateUsername(username); err != nil {
		return nil, err
	}
	if err := ValidateFirstName(firstName); err != nil {
		return nil, err
	}
	if err := ValidateLastName(lastName); err != nil {
		return nil, err
	}
	if err := ValidateEmail(email); err != nil {
		return nil, err
	}
	if err := validatePassword(password); err != nil {
		return nil, err
	}
	if err := ValidatePhoneNumber(phoneNumber); err != nil {
		return nil, err
	}

	hashedPassword, err := hashPassword(password)
	if err != nil {
		return nil, err
	}

	return &User{
		Username:      username,
		FirstName:     firstName,
		LastName:      lastName,
		Email:         email,
		Password:      hashedPassword,
		IsActive:      true,
		Role:          "user",
		PhoneNumber:   phoneNumber,
		EmailVerified: false,
		CreatedAt:     time.Now(),
		UpdatedAt:     time.Now(),
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
	u.UpdatedAt = time.Now()
	return nil
}

func (u *User) UpdateProfile(firstName, lastName, phoneNumber string) error {
	if err := ValidateFirstName(firstName); err != nil {
		return err
	}
	if err := ValidateLastName(lastName); err != nil {
		return err
	}
	if err := ValidatePhoneNumber(phoneNumber); err != nil {
		return err
	}

	u.FirstName = firstName
	u.LastName = lastName
	u.PhoneNumber = phoneNumber
	u.UpdatedAt = time.Now()
	return nil
}

func (u *User) UpdateEmailVerification(verified bool) {
	u.EmailVerified = verified
	u.UpdatedAt = time.Now()
}

func (u *User) UpdateLastLogin() {
	u.LastLogin = time.Now()
	u.UpdatedAt = time.Now()
}

func (u *User) SetProfileImageURL(url string) {
	u.ProfileImageURL = url
	u.UpdatedAt = time.Now()
}

func (u *User) SetRole(role string) error {
	validRoles := []string{"user", "admin", "moderator"}
	for _, validRole := range validRoles {
		if role == validRole {
			u.Role = role
			u.UpdatedAt = time.Now()
			return nil
		}
	}
	return errors.New("invalid role")
}

func ValidateUsername(username string) error {
	if len(username) < 3 || len(username) > 50 {
		return errors.New("username must be between 3 and 50 characters")
	}
	return nil
}

func ValidateFirstName(firstName string) error {
	if len(firstName) == 0 || len(firstName) > 50 {
		return errors.New("first name must be between 1 and 50 characters")
	}
	return nil
}

func ValidateLastName(lastName string) error {
	if len(lastName) == 0 || len(lastName) > 50 {
		return errors.New("last name must be between 1 and 50 characters")
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

func ValidatePhoneNumber(phoneNumber string) error {
	phoneRegex := regexp.MustCompile(`^\+?[1-9]\d{1,14}$`)
	if !phoneRegex.MatchString(phoneNumber) {
		return errors.New("invalid phone number format")
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
