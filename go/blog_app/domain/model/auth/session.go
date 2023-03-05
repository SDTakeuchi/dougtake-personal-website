package auth

import (
	"blog_app/domain/model/uuid"
	"time"
)

type Session interface {
	// ID originally comes from refresh token's id
	ID() uuid.UUID
	UserID() uuid.UUID
	RefreshToken() string
	UserAgent() string
	ClientIp() string
	// IsBlocked()    bool
	ExpiresAt() time.Time
	CreatedAt() time.Time
}

// ID           uuid.UUID `json:"id"`
// Username     string    `json:"username"`
// RefreshToken string    `json:"refresh_token"`
// UserAgent    string    `json:"user_agent"`
// ClientIp     string    `json:"client_ip"`
// IsBlocked    bool      `json:"is_blocked"`
// ExpiresAt    time.Time `json:"expires_at"`
// CreatedAt    time.Time `json:"created_at"`
