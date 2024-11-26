// File: internal/application/health/health_service.go
package health

import (
	"context"
	"quickflow/internal/domain/health"
	"quickflow/pkg/errors"
	"time"
)

type DatabaseChecker interface {
	CheckConnection(ctx context.Context) error
}

type HealthService struct {
	db DatabaseChecker
}

func NewHealthService(db DatabaseChecker) *HealthService {
	return &HealthService{
		db: db,
	}
}

func (s *HealthService) CheckHealth(ctx context.Context) (*health.SystemHealth, error) {
	// Create context with timeout for database check
	ctxWithTimeout, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	// Check database health
	dbStatus := health.Status{
		Status:  "ok",
		Message: "Database is responding",
	}

	if err := s.db.CheckConnection(ctxWithTimeout); err != nil {
		dbStatus = health.Status{
			Status:  "error",
			Message: "Database connection failed",
		}
		return &health.SystemHealth{
			Server: health.Status{
				Status:  "ok",
				Message: "Server is running",
			},
			Database: dbStatus,
		}, errors.NewAppError(errors.ErrorTypeInternal, "Database health check failed", err)
	}

	return &health.SystemHealth{
		Server: health.Status{
			Status:  "ok",
			Message: "Server is running",
		},
		Database: dbStatus,
	}, nil
}
