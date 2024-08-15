// File: internal/infrastructure/repository/user_repository.go

package repository

import (
	"quickflow/internal/domain/user"

	"gorm.io/gorm"
)

type UserRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) *UserRepository {
	return &UserRepository{db: db}
}

func (r *UserRepository) Create(user *user.User) error {
	return r.db.Create(user).Error
}

func (r *UserRepository) GetByID(id uint) (*user.User, error) {
	var user user.User
	err := r.db.First(&user, id).Error
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func (r *UserRepository) GetByEmail(email string) (*user.User, error) {
	var user user.User
	err := r.db.Where("email = ?", email).First(&user).Error
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func (r *UserRepository) Update(user *user.User) error {
	return r.db.Save(user).Error
}

func (r *UserRepository) Delete(id uint) error {
	return r.db.Delete(&user.User{}, id).Error
}
