// File: internal/infrastructure/repository/health_repository.go
package repository

import (
	"context"
	"quickflow/pkg/errors"

	"gorm.io/gorm"
)

type HealthRepository struct {
	db *gorm.DB
}

func NewHealthRepository(db *gorm.DB) *HealthRepository {
	return &HealthRepository{
		db: db,
	}
}

func (r *HealthRepository) CheckConnection(ctx context.Context) error {
	sqlDB, err := r.db.DB()
	if err != nil {
		return errors.Wrap(err, "failed to get database instance")
	}

	if err := sqlDB.PingContext(ctx); err != nil {
		return errors.Wrap(err, "failed to ping database")
	}

	return nil
}
