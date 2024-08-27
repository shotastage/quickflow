// File: internal/infrastructure/repository/loginhistory_repository.go

package repository

import (
	"context"

	"quickflow/internal/domain/loginhistory"

	"gorm.io/gorm"
)

type LoginHistoryRepository struct {
	db *gorm.DB
}

func NewLoginHistoryRepository(db *gorm.DB) *LoginHistoryRepository {
	return &LoginHistoryRepository{db: db}
}

func (r *LoginHistoryRepository) Create(ctx context.Context, history *loginhistory.LoginHistory) error {
	return r.db.WithContext(ctx).Create(history).Error
}

func (r *LoginHistoryRepository) GetByUserID(ctx context.Context, userID string) ([]*loginhistory.LoginHistory, error) {
	var histories []*loginhistory.LoginHistory
	err := r.db.WithContext(ctx).Where("user_id = ?", userID).Order("login_time DESC").Find(&histories).Error
	return histories, err
}
