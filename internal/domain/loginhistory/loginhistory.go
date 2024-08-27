// File: internal/domain/loginhistory/loginhistory.go

package loginhistory

import (
	"time"

	"github.com/google/uuid"
)

type LoginHistory struct {
	ID        uuid.UUID `json:"id"`
	UserID    uuid.UUID `json:"user_id"`
	IPAddress string    `json:"ip_address"`
	UserAgent string    `json:"user_agent"`
	LoginTime time.Time `json:"login_time"`
}

func NewLoginHistory(userID uuid.UUID, ipAddress, userAgent string) *LoginHistory {
	return &LoginHistory{
		ID:        uuid.New(),
		UserID:    userID,
		IPAddress: ipAddress,
		UserAgent: userAgent,
		LoginTime: time.Now(),
	}
}
