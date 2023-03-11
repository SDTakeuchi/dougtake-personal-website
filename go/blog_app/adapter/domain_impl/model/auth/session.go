package auth

import (
	"blog_app/domain/model/auth"
	"blog_app/domain/model/uuid"
	"context"
	"time"
)

type session struct {
	id           uuid.UUID
	userID       uuid.UUID
	refreshToken string
	userAgent    string
	clientIP     string
	expiresAt    time.Time
	createdAt    time.Time
}

func (s *session) ID() uuid.UUID { return s.id }

func (s *session) UserID() uuid.UUID { return s.userID }

func (s *session) RefreshToken() string { return s.refreshToken }

func (s *session) UserAgent() string { return s.userAgent }

func (s *session) ClientIP() string { return s.clientIP }

func (s *session) ExpiresAt() time.Time { return s.expiresAt }

func (s *session) CreatedAt() time.Time { return s.createdAt }

func NewSession(
	id uuid.UUID,
	userID uuid.UUID,
	refreshToken string,
	userAgent string,
	clientIP string,
	expiresAt time.Time,
	createdAt time.Time,
) (auth.Session, error) {
	s := &session{
		id:           id,
		userID:       userID,
		refreshToken: refreshToken,
		userAgent:    userAgent,
		clientIP:     clientIP,
		expiresAt:    expiresAt,
		createdAt:    createdAt,
	}
	if err := auth.ValidateSession(s); err != nil {
		return nil, err
	}
	return s, nil
}
