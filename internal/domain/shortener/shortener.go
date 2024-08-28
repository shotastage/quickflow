// File: internal/domain/shortener/shortener.go

package shortener

import (
	"time"
)

type ShortenedURL struct {
	ID             uint   `gorm:"primaryKey"`
	Code           string `gorm:"uniqueIndex;not null"`
	OriginalURL    string `gorm:"not null"`
	IsBanned       bool
	SecurityReason string
	CreatedAt      time.Time
	UpdatedAt      time.Time
}
