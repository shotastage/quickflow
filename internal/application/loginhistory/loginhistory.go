// File: internal/application/loginhistory/loginhistory_service.go

package loginhistory

import (
	"context"

	"quickflow/internal/domain/loginhistory"
)

type LoginHistoryRepository interface {
	Create(ctx context.Context, history *loginhistory.LoginHistory) error
	GetByUserID(ctx context.Context, userID string) ([]*loginhistory.LoginHistory, error)
}

type LoginHistoryService struct {
	repo LoginHistoryRepository
}

func NewLoginHistoryService(repo LoginHistoryRepository) *LoginHistoryService {
	return &LoginHistoryService{repo: repo}
}

func (s *LoginHistoryService) RecordLogin(ctx context.Context, history *loginhistory.LoginHistory) error {
	return s.repo.Create(ctx, history)
}

func (s *LoginHistoryService) GetUserLoginHistory(ctx context.Context, userID string) ([]*loginhistory.LoginHistory, error) {
	return s.repo.GetByUserID(ctx, userID)
}
