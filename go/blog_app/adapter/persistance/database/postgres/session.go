package postgres

import (
	"time"
)

type Session struct {
	ID           string `json:"id"`
	Username     string `json:"username"`
	RefreshToken string `json:"refresh_token"`
	UserAgent    string `json:"user_agent"`
	ClientIp     string `json:"client_ip"`
	// IsBlocked    bool      `json:"is_blocked"`
	ExpiresAt time.Time `json:"expires_at"`
}
