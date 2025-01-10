// File: internal/domain/sqlexecutor/sql_executor.go

package sqlexecutor

import (
	"time"
)

// QueryRequest represents a SQL query execution request
type QueryRequest struct {
	Query   string                 `json:"query"`
	Params  map[string]interface{} `json:"params,omitempty"`
	Timeout time.Duration          `json:"timeout,omitempty"`
}

// QueryResult represents the result of a SQL query execution
type QueryResult struct {
	Columns       []string        `json:"columns"`
	Rows          [][]interface{} `json:"rows"`
	RowCount      int64           `json:"rowCount"`
	ExecutionTime time.Duration   `json:"executionTime"`
}

// ValidationError represents an error during query validation
type ValidationError struct {
	Message string `json:"message"`
	Query   string `json:"query"`
}

// Error implements the error interface
func (e *ValidationError) Error() string {
	return e.Message
}
