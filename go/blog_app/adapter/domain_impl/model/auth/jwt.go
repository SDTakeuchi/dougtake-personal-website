package auth

import (
	"errors"
	"fmt"
	"time"

	"blog_app/adapter/config"
	"blog_app/domain/model/auth"
	"blog_app/domain/model/uuid"

	"github.com/golang-jwt/jwt/v5"
)

var jwtMethod = jwt.SigningMethodHS256

type JWTIssuer struct {
	secretKey string
}

func NewJWTIssuer(secretKey string) (auth.TokenIssuer, error) {
	minSecretKeySize := config.Get().Token.MinSecretKeySize

	if len(secretKey) < minSecretKeySize {
		return nil, fmt.Errorf("invalid key size: must be at least %d characters", minSecretKeySize)
	}
	return &JWTIssuer{secretKey}, nil
}

func (issuer *JWTIssuer) Create(
	userID uuid.UUID,
	duration time.Duration,
) (string, auth.Payload, error) {
	p := NewPayload(userID, duration)
	claims := ClaimsFromPayload(p)
	token := jwt.NewWithClaims(jwtMethod, claims)
	tokenString, err := token.SignedString([]byte(issuer.secretKey))
	return tokenString, p, err
}

func (issuer *JWTIssuer) Verify(token string) (auth.Payload, error) {
	keyFunc := func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, auth.ErrInvalidToken
		}
		return []byte(issuer.secretKey), nil
	}

	parsedToken, err := jwt.Parse(token, keyFunc)
	if err != nil {
		if errors.Is(err, jwt.ErrTokenExpired) {
			return nil, auth.ErrExpiredToken
		}
		return nil, auth.ErrInvalidToken
	}
	if !parsedToken.Valid {
		return nil, auth.ErrInvalidToken
	}

	claims, ok := parsedToken.Claims.(jwt.MapClaims)
	if !ok {
		return nil, auth.ErrInvalidToken
	}

	payload, err := PayloadFromClaims(claims)
	if err != nil {
		return nil, err
	}
	return payload, nil
}
