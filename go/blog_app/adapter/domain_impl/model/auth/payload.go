package auth

import (
	"fmt"
	"time"

	"blog_app/domain/model/auth"
	"blog_app/domain/model/uuid"

	"github.com/golang-jwt/jwt/v5"
)

type payload struct {
	id        uuid.UUID
	userID    uuid.UUID
	issuedAt  time.Time
	expiresAt time.Time
}

func (p *payload) ID() uuid.UUID { return p.id }

func (p *payload) UserID() uuid.UUID { return p.userID }

func (p *payload) IssuedAt() time.Time { return p.issuedAt }

func (p *payload) ExpiresAt() time.Time { return p.expiresAt }

func NewPayload(userID uuid.UUID, duration time.Duration) auth.Payload {
	return &payload{
		id:        uuid.New(),
		userID:    userID,
		issuedAt:  time.Now(),
		expiresAt: time.Now().Add(duration),
	}
}

func ClaimsFromPayload(p auth.Payload) jwt.MapClaims {
	return jwt.MapClaims{
		"id":  p.ID(),
		"sub": p.UserID(),
		"exp": p.ExpiresAt().Unix(),
		"iat": p.IssuedAt().Unix(),
	}
}

func PayloadFromClaims(claims jwt.MapClaims) (auth.Payload, error) {
	payloadID, ok := claims["id"]
	if !ok {
		return nil, fmt.Errorf("payload id did not match")
	}
	payloadIDString, ok := payloadID.(string)
	if !ok {
		return nil, fmt.Errorf("invalid payload id")
	}
	payloadIDUUID, err := uuid.Parse(payloadIDString)
	if err != nil {
		return nil, err
	}
	userID, ok := claims["sub"]
	if !ok {
		return nil, fmt.Errorf("user id did not match")
	}
	id, ok := userID.(string)
	if !ok {
		return nil, fmt.Errorf("user id did not match")
	}
	parsedID, err := uuid.Parse(id)
	if err != nil {
		return nil, err
	}
	issuedAt, err := claims.GetIssuedAt()
	if err != nil {
		return nil, err
	}
	expiresAt, err := claims.GetExpirationTime()
	if err != nil {
		return nil, err
	}
	return &payload{
		id:        payloadIDUUID,
		userID:    parsedID,
		issuedAt:  issuedAt.Time,
		expiresAt: expiresAt.Time,
	}, nil
}
