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

type TokenType int

const (
	AccessToken TokenType = iota + 1
	RefreshToken
)

type Payload interface {
	ID() uuid.UUID
	TokenType() TokenType
	UserID() uuid.UUID
	IssuedAt() time.Time
	ExpiresAt() time.Time
}
