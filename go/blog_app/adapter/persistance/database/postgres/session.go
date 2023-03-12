package postgres

import (
	"time"
)

type Session struct {
	ID           string `json:"id"`
	UserID       string `json:"user_id"`
	RefreshToken string `json:"refresh_token"`
	UserAgent    string `json:"user_agent"`
	ClientIP     string `json:"client_ip"`
	ExpiresAt time.Time `json:"expires_at"`
	CreatedAt time.Time `json:"created_at"`
}
