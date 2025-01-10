// File: internal/application/sqlservice/sql_executor_service.go

package sqlservice

import (
	"context"
	"database/sql"
	"quickflow/internal/domain/sqlexecutor"
	"quickflow/pkg/errors"
	"time"
)

type SQLExecutorService interface {
	ExecuteQuery(ctx context.Context, req sqlexecutor.QueryRequest) (*sqlexecutor.QueryResult, error)
}

type sqlExecutorService struct {
	db        *sql.DB
	validator QueryValidator
}

func NewSQLExecutorService(db *sql.DB) SQLExecutorService {
	return &sqlExecutorService{
		db:        db,
		validator: NewQueryValidator(),
	}
}

func (s *sqlExecutorService) ExecuteQuery(ctx context.Context, req sqlexecutor.QueryRequest) (*sqlexecutor.QueryResult, error) {
	// Validate query
	if err := s.validator.Validate(req.Query); err != nil {
		return nil, errors.NewAppError(errors.ErrorTypeValidation, "Invalid query", err)
	}

	// Set timeout if specified
	if req.Timeout > 0 {
		var cancel context.CancelFunc
		ctx, cancel = context.WithTimeout(ctx, req.Timeout)
		defer cancel()
	}

	startTime := time.Now()

	// Execute query
	rows, err := s.db.QueryContext(ctx, req.Query)
	if err != nil {
		return nil, errors.NewAppError(errors.ErrorTypeInternal, "Query execution failed", err)
	}
	defer rows.Close()

	// Get column names
	columns, err := rows.Columns()
	if err != nil {
		return nil, errors.NewAppError(errors.ErrorTypeInternal, "Failed to get column names", err)
	}

	// Prepare result
	var result [][]interface{}
	for rows.Next() {
		// Create a slice of interface{} to hold the values
		values := make([]interface{}, len(columns))
		scanArgs := make([]interface{}, len(columns))
		for i := range values {
			scanArgs[i] = &values[i]
		}

		// Scan the row into the values slice
		if err := rows.Scan(scanArgs...); err != nil {
			return nil, errors.NewAppError(errors.ErrorTypeInternal, "Failed to scan row", err)
		}

		result = append(result, values)
	}

	if err = rows.Err(); err != nil {
		return nil, errors.NewAppError(errors.ErrorTypeInternal, "Error during row iteration", err)
	}

	return &sqlexecutor.QueryResult{
		Columns:       columns,
		Rows:          result,
		RowCount:      int64(len(result)),
		ExecutionTime: time.Since(startTime),
	}, nil
}
