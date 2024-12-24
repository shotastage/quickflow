// File: internal/application/table/table_service.go
package table

import (
	"context"
	"fmt"
	"quickflow/internal/domain/tableentity"
	"quickflow/pkg/errors"

	"regexp"
)

type TableRepository interface {
	CreateTable(ctx context.Context, table *tableentity.Table) error
	TableExists(ctx context.Context, tableName string) (bool, error)
}

type TableService struct {
	repo TableRepository
}

func NewTableService(repo TableRepository) *TableService {
	return &TableService{repo: repo}
}

func (s *TableService) CreateTable(ctx context.Context, table *tableentity.Table) error {
	if err := validateTable(table); err != nil {
		return err
	}

	exists, err := s.repo.TableExists(ctx, table.Name)
	if err != nil {
		return err
	}
	if exists {
		return errors.NewAppError(
			errors.ErrorTypeValidation,
			fmt.Sprintf("Table '%s' already exists", table.Name),
			nil,
		)
	}

	return s.repo.CreateTable(ctx, table)
}

func validateTable(table *tableentity.Table) error {
	// テーブル名のバリデーション
	if table.Name == "" {
		return errors.NewAppError(
			errors.ErrorTypeValidation,
			"Table name is required",
			nil,
		)
	}

	// テーブル名は英数字とアンダースコアのみ許可
	if !regexp.MustCompile(`^[a-zA-Z][a-zA-Z0-9_]*$`).MatchString(table.Name) {
		return errors.NewAppError(
			errors.ErrorTypeValidation,
			"Table name must start with a letter and contain only letters, numbers, and underscores",
			nil,
		)
	}

	if len(table.Columns) == 0 {
		return errors.NewAppError(
			errors.ErrorTypeValidation,
			"At least one column is required",
			nil,
		)
	}

	// カラムのバリデーション
	columnNames := make(map[string]bool)
	hasPrimaryKey := false

	for _, col := range table.Columns {
		if col.Name == "" {
			return errors.NewAppError(
				errors.ErrorTypeValidation,
				"Column name is required",
				nil,
			)
		}

		if !regexp.MustCompile(`^[a-zA-Z][a-zA-Z0-9_]*$`).MatchString(col.Name) {
			return errors.NewAppError(
				errors.ErrorTypeValidation,
				fmt.Sprintf("Invalid column name: %s", col.Name),
				nil,
			)
		}

		if columnNames[col.Name] {
			return errors.NewAppError(
				errors.ErrorTypeValidation,
				fmt.Sprintf("Duplicate column name: %s", col.Name),
				nil,
			)
		}
		columnNames[col.Name] = true

		if col.PrimaryKey {
			hasPrimaryKey = true
		}
	}

	if !hasPrimaryKey {
		return errors.NewAppError(
			errors.ErrorTypeValidation,
			"Table must have at least one primary key column",
			nil,
		)
	}

	return nil
}
