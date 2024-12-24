// File: internal/infrastructure/repository/table_repository.go
package repository

import (
	"context"
	"fmt"
	"quickflow/internal/domain/tableentity"
	"quickflow/pkg/errors"
	"strings"

	"gorm.io/gorm"
)

type TableRepository struct {
	db *gorm.DB
}

func NewTableRepository(db *gorm.DB) *TableRepository {
	return &TableRepository{db: db}
}

func (r *TableRepository) TableExists(ctx context.Context, tableName string) (bool, error) {
	var exists bool
	err := r.db.WithContext(ctx).Raw(
		"SELECT EXISTS (SELECT FROM information_schema.tables WHERE table_name = ?)",
		tableName,
	).Scan(&exists).Error

	if err != nil {
		return false, errors.NewAppError(
			errors.ErrorTypeInternal,
			"Failed to check table existence",
			err,
		)
	}
	return exists, nil
}

func (r *TableRepository) CreateTable(ctx context.Context, table *tableentity.Table) error {
	sql := buildCreateTableSQL(table)

	if err := r.db.WithContext(ctx).Exec(sql).Error; err != nil {
		return errors.NewAppError(
			errors.ErrorTypeInternal,
			"Failed to create table",
			err,
		)
	}

	// テーブルにコメントが指定されている場合
	if table.Description != "" {
		commentSQL := fmt.Sprintf(
			"COMMENT ON TABLE %s IS '%s'",
			table.Name,
			strings.Replace(table.Description, "'", "''", -1),
		)
		if err := r.db.WithContext(ctx).Exec(commentSQL).Error; err != nil {
			return errors.NewAppError(
				errors.ErrorTypeInternal,
				"Failed to add table comment",
				err,
			)
		}
	}

	return nil
}

func buildCreateTableSQL(table *tableentity.Table) string {
	var columnDefs []string

	for _, col := range table.Columns {
		def := fmt.Sprintf("%s %s", col.Name, col.Type)

		if col.Type == tableentity.TypeVARCHAR && col.Length != nil {
			def = fmt.Sprintf("%s(%d)", def, *col.Length)
		}

		if col.NotNull {
			def += " NOT NULL"
		}

		if col.PrimaryKey {
			def += " PRIMARY KEY"
			if col.AutoIncrement {
				if col.Type == tableentity.TypeINT {
					def = fmt.Sprintf("%s SERIAL PRIMARY KEY", col.Name)
				} else if col.Type == tableentity.TypeBIGINT {
					def = fmt.Sprintf("%s BIGSERIAL PRIMARY KEY", col.Name)
				}
			}
		}

		if col.Unique {
			def += " UNIQUE"
		}

		if col.Default != nil {
			def += fmt.Sprintf(" DEFAULT %s", *col.Default)
		}

		columnDefs = append(columnDefs, def)
	}

	return fmt.Sprintf(
		"CREATE TABLE %s (\n  %s\n)",
		table.Name,
		strings.Join(columnDefs, ",\n  "),
	)
}
