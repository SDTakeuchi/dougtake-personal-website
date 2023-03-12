package auth

import (
	"errors"
	"fmt"
	"time"

	"blog_app/adapter/config"
	"blog_app/domain/model/auth"
	"blog_app/domain/model/uuid"

	"github.com/golang-jwt/jwt/v5"
	"github.com/pingcap/log"
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
	tokenType auth.TokenType,
	duration time.Duration,
) (string, auth.Payload, error) {
	// test code cannot read config, so provide default value here
	if duration == time.Duration(0) {
		duration = time.Minute
	}
	p := NewPayload(userID, tokenType, duration)
	claims := ClaimsFromPayload(p)
	token := jwt.NewWithClaims(jwtMethod, claims)
	signedToken, err := token.SignedString([]byte(issuer.secretKey))
	return signedToken, p, err
}

func (issuer *JWTIssuer) Verify(token string) (auth.Payload, error) {
	keyFunc := func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			log.Info("JWTIssuer.Veriy failed to assert")
			return nil, auth.ErrInvalidToken
		}
		return []byte(issuer.secretKey), nil
	}

	parsedToken, err := jwt.Parse(token, keyFunc)
	if err != nil {
		if errors.Is(err, jwt.ErrTokenExpired) {
			return nil, auth.ErrExpiredToken
		}
		log.Info(fmt.Sprintf("JWTIssuer.Veriy failed: %v", err.Error()))
		return nil, auth.ErrInvalidToken
	}
	if !parsedToken.Valid {
		log.Info(fmt.Sprintf("JWTIssuer.Veriy failed: %v", err.Error()))
		return nil, auth.ErrInvalidToken
	}

	claims, ok := parsedToken.Claims.(jwt.MapClaims)
	if !ok {
		log.Info("failed: parsedToken.Claims.(jwt.MapClaims)")
		return nil, auth.ErrInvalidToken
	}

	payload, err := PayloadFromClaims(claims)
	if err != nil {
		log.Info(fmt.Sprintf("JWTIssuer.Veriy failed: %v", err.Error()))
		return nil, err
	}
	return payload, nil
}
