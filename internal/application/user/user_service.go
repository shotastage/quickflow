// File: internal/application/user/user_service.go

package user

import (
	"errors"

	"quickflow/internal/domain/user"
)

type UserRepository interface {
	Create(user *user.User) error
	GetByID(id uint) (*user.User, error)
	GetByEmail(email string) (*user.User, error)
	Update(user *user.User) error
	Delete(id uint) error
}

type UserService struct {
	repo UserRepository
}

func NewUserService(repo UserRepository) *UserService {
	return &UserService{repo: repo}
}

func (s *UserService) CreateUser(username, firstName, lastName, email, password, phoneNumber string) (*user.User, error) {
	newUser, err := user.NewUser(username, firstName, lastName, email, password, phoneNumber)
	if err != nil {
		return nil, err
	}

	if err := s.repo.Create(newUser); err != nil {
		return nil, err
	}

	return newUser, nil
}

func (s *UserService) GetUserByID(id uint) (*user.User, error) {
	return s.repo.GetByID(id)
}

func (s *UserService) GetUserByEmail(email string) (*user.User, error) {
	return s.repo.GetByEmail(email)
}

func (s *UserService) UpdateUser(id uint, username, firstName, lastName, email, phoneNumber string) (*user.User, error) {
	existingUser, err := s.repo.GetByID(id)
	if err != nil {
		return nil, err
	}

	if username != "" {
		if err := user.ValidateUsername(username); err != nil {
			return nil, err
		}
		existingUser.Username = username
	}

	if firstName != "" {
		if err := user.ValidateFirstName(firstName); err != nil {
			return nil, err
		}
		existingUser.FirstName = firstName
	}

	if lastName != "" {
		if err := user.ValidateLastName(lastName); err != nil {
			return nil, err
		}
		existingUser.LastName = lastName
	}

	if email != "" {
		if err := user.ValidateEmail(email); err != nil {
			return nil, err
		}
		existingUser.Email = email
	}

	if phoneNumber != "" {
		if err := user.ValidatePhoneNumber(phoneNumber); err != nil {
			return nil, err
		}
		existingUser.PhoneNumber = phoneNumber
	}

	if err := s.repo.Update(existingUser); err != nil {
		return nil, err
	}

	return existingUser, nil
}

func (s *UserService) UpdatePassword(id uint, oldPassword, newPassword string) error {
	existingUser, err := s.repo.GetByID(id)
	if err != nil {
		return err
	}

	if !user.CheckPasswordHash(oldPassword, existingUser.Password) {
		return errors.New("old password is incorrect")
	}

	if err := existingUser.UpdatePassword(newPassword); err != nil {
		return err
	}

	return s.repo.Update(existingUser)
}

func (s *UserService) UpdateProfileImage(id uint, imageURL string) error {
	existingUser, err := s.repo.GetByID(id)
	if err != nil {
		return err
	}

	existingUser.SetProfileImageURL(imageURL)

	return s.repo.Update(existingUser)
}

func (s *UserService) UpdateRole(id uint, role string) error {
	existingUser, err := s.repo.GetByID(id)
	if err != nil {
		return err
	}

	if err := existingUser.SetRole(role); err != nil {
		return err
	}

	return s.repo.Update(existingUser)
}

func (s *UserService) VerifyEmail(id uint) error {
	existingUser, err := s.repo.GetByID(id)
	if err != nil {
		return err
	}

	existingUser.UpdateEmailVerification(true)

	return s.repo.Update(existingUser)
}

func (s *UserService) UpdateLastLogin(id uint) error {
	existingUser, err := s.repo.GetByID(id)
	if err != nil {
		return err
	}

	existingUser.UpdateLastLogin()

	return s.repo.Update(existingUser)
}

func (s *UserService) DeleteUser(id uint) error {
	return s.repo.Delete(id)
}
