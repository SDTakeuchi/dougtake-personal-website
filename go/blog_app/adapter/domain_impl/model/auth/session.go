package auth

import (
	"blog_app/adapter/persistance/database/postgres"
	"blog_app/domain/model/auth"
	"blog_app/domain/model/uuid"
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

func SessionFromRecord(r postgres.Session) auth.Session {
	// assume ids are always parsable
	id, _ := uuid.Parse(r.ID)
	userIDd, _ := uuid.Parse(r.UserID)
	return &session{
		id:           id,
		userID:       userIDd,
		refreshToken: r.RefreshToken,
		userAgent:    r.UserAgent,
		clientIP:     r.ClientIP,
		expiresAt:    r.ExpiresAt,
		createdAt:    r.CreatedAt,
	}
}

func SessionToRecord(m auth.Session) postgres.Session {
	return postgres.Session{
		ID:           m.ID().String(),
		UserID:       m.UserID().String(),
		RefreshToken: m.RefreshToken(),
		UserAgent:    m.UserAgent(),
		ClientIP:     m.ClientIP(),
		ExpiresAt:    m.ExpiresAt(),
		CreatedAt:    m.CreatedAt(),
	}
}
