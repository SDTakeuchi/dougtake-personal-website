package usecase

import (
	"blog_app/adapter/config"
	"blog_app/domain/model/auth"
	"blog_app/domain/repository"
	"context"
	"fmt"
	"time"
)

type (
	RenewToken interface {
		Execute(ctx context.Context, input RenewTokenInput) (*RenewTokenOutput, error)
	}
	RenewTokenInput struct {
		Token string
	}
	RenewTokenOutput struct {
		AccessToken          string
		AccessTokenExpiresAt time.Time
	}
	renewTokenImpl struct {
		tokenIssuer auth.TokenIssuer
		sessionRepo repository.Session
	}
)

func NewRenewToken(
	tokenIssuer auth.TokenIssuer,
	sessionRepo repository.Session,
) RenewToken {
	return &renewTokenImpl{
		tokenIssuer: tokenIssuer,
		sessionRepo: sessionRepo,
	}
}

func (u *renewTokenImpl) Execute(ctx context.Context, input RenewTokenInput) (*RenewTokenOutput, error) {
	payload, err := u.tokenIssuer.Verify(input.Token)
	if err != nil {
		return nil, err
	}

	if payload.TokenType() != auth.RefreshToken {
		return nil, fmt.Errorf(
			"%w: use refresh token to renew token",
			auth.ErrInvalidToken,
		)
	}

	session, err := u.sessionRepo.Get(ctx, payload.ID().String())
	if err != nil {
		return nil, fmt.Errorf(
			"%w: %s",
			auth.ErrInvalidToken,
			err.Error(),
		)
	}

	if payload.UserID() != session.UserID() {
		return nil, fmt.Errorf(
			"%w: user id not correct",
			auth.ErrInvalidToken,
		)
	}

	if input.Token != session.RefreshToken() {
		return nil, fmt.Errorf(
			"%w: session token not correct",
			auth.ErrInvalidToken,
		)
	}

	if time.Now().After(session.ExpiresAt()) {
		return nil, auth.ErrExpiredToken
	}

	accessToken, accessPayload, err := u.tokenIssuer.Create(
		payload.UserID(),
		auth.AccessToken,
		config.Get().Token.AccessTokenDuration,
	)
	if err != nil {
		return nil, err
	}

	return &RenewTokenOutput{
		AccessToken:          accessToken,
		AccessTokenExpiresAt: accessPayload.ExpiresAt(),
	}, nil
}
