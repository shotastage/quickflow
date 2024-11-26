// File: internal/domain/health/health.go
package health

// Status represents the health status of a system component
type Status struct {
	Status  string `json:"status"`
	Message string `json:"message"`
}

// SystemHealth represents the overall system health status
type SystemHealth struct {
	Server   Status `json:"server"`
	Database Status `json:"database"`
}
