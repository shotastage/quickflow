// File: internal/domain/query/query.go
package query

type Query struct {
	ID        string `json:"id"`
	SQL       string `json:"sql"`
	CreatedAt int64  `json:"created_at"`
	UpdatedAt int64  `json:"updated_at"`
}

type QueryResult struct {
	Columns []string        `json:"columns"`
	Rows    [][]interface{} `json:"rows"`
}

type QueryExecutionError struct {
	Message string `json:"message"`
	Code    string `json:"code"`
}
