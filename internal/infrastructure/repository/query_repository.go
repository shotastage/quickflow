// File: internal/domain/query/repository.go
package repository

type Repository interface {
	ExecuteQuery(sql string) (*QueryResult, error)
	ValidateQuery(sql string) error
	SaveQuery(query *Query) error
	GetQueryHistory() ([]*Query, error)
}
