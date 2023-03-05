package auth

import (
	"errors"
	"time"

	"blog_app/domain/model/uuid"
)

var (
	// errors returned by VerifyToken method
	ErrExpiredToken = errors.New("token is expired")
	ErrInvalidToken = errors.New("token is invalid")
)

type Payload interface {
	ID() uuid.UUID
	UserID() uuid.UUID
	IssuedAt() time.Time
	ExpiresAt() time.Time
}
