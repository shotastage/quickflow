// File: internal/application/loginhistory/loginhistory_service.go

package loginhistory

import (
	"context"

	"quickflow/internal/domain/loginhistory"

	"github.com/google/uuid"
)

type Repository interface {
	Create(ctx context.Context, history *loginhistory.LoginHistory) error
	GetByUserID(ctx context.Context, userID uuid.UUID) ([]*loginhistory.LoginHistory, error)
}

type Service interface {
	RecordLogin(ctx context.Context, userID uuid.UUID, ipAddress, userAgent string) error
	GetUserLoginHistory(ctx context.Context, userID uuid.UUID) ([]*loginhistory.LoginHistory, error)
}

type loginHistoryService struct {
	repo Repository
}

func NewLoginHistoryService(repo Repository) Service {
	return &loginHistoryService{repo: repo}
}

func (s *loginHistoryService) RecordLogin(ctx context.Context, userID uuid.UUID, ipAddress, userAgent string) error {
	history := loginhistory.NewLoginHistory(userID, ipAddress, userAgent)
	return s.repo.Create(ctx, history)
}

func (s *loginHistoryService) GetUserLoginHistory(ctx context.Context, userID uuid.UUID) ([]*loginhistory.LoginHistory, error) {
	return s.repo.GetByUserID(ctx, userID)
}
