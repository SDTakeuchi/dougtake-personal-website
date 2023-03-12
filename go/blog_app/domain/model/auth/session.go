package auth

import (
	"blog_app/domain/model"
	"blog_app/domain/model/uuid"
	"fmt"
	"time"
)

type Session interface {
	// ID originally comes from refresh token's id
	ID() uuid.UUID
	UserID() uuid.UUID
	RefreshToken() string
	UserAgent() string
	ClientIP() string
	// IsBlocked()    bool
	ExpiresAt() time.Time
	CreatedAt() time.Time
}

func ValidateSession(s Session) error {
	if s.RefreshToken() == "" {
		return fmt.Errorf(
			"%w: session must have a refresh token string",
			model.ErrConstructor,
		)
	}
	if s.ExpiresAt().IsZero() {
		return fmt.Errorf(
			"%w: session must have expiration date",
			model.ErrConstructor,
		)
	}
	return nil
}
