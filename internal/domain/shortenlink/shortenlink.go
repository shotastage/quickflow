// File: internal/domain/shortener/shortener.go

package shortenlink

import (
	"time"
)

type ShortenLink struct {
	ID             uint      `gorm:"primaryKey;autoIncrement"`
	Code           string    `gorm:"uniqueIndex;not null;size:255"`
	OriginalURL    string    `gorm:"not null;size:2048"`
	IsBanned       bool      `gorm:"default:false"`
	SecurityReason string    `gorm:"size:255"`
	CreatedAt      time.Time `gorm:"autoCreateTime"`
	UpdatedAt      time.Time `gorm:"autoUpdateTime"`
}
