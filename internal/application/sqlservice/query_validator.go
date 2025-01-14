// File: internal/application/sqlservice/query_validator.go

package sqlservice

import (
	"quickflow/internal/domain/sqlexecutor"
	"strings"
)

type QueryValidator interface {
	Validate(query string) error
}

type queryValidator struct {
	forbiddenKeywords []string
}

func NewQueryValidator() QueryValidator {
	return &queryValidator{
		forbiddenKeywords: []string{
			"DROP", "TRUNCATE", "DELETE", "ALTER", "CREATE", "INSERT",
			"UPDATE", "GRANT", "REVOKE",
		},
	}
}

func (v *queryValidator) Validate(query string) error {
	upperQuery := strings.ToUpper(query)

	// Check for forbidden keywords
	for _, keyword := range v.forbiddenKeywords {
		if strings.Contains(upperQuery, keyword) {
			return &sqlexecutor.ValidationError{
				Message: "Query contains forbidden keyword: " + keyword,
				Query:   query,
			}
		}
	}

	return nil
}
