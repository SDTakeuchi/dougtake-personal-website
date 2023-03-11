package usecase

import (
	"blog_app/adapter/config"
	"blog_app/domain/model/auth"
	"blog_app/domain/model/password"
	"blog_app/domain/model/uuid"
	"blog_app/domain/repository"
	"context"
	"time"
)

type (
	RenewToken interface {
		Execute(ctx context.Context, input RenewTokenInput) (*RenewTokenOutput, error)
	}
	RenewTokenInput struct {
		token string
	}
	RenewTokenOutput struct {
		AccessToken           string
		AccessTokenExpiresAt  time.Time
		RefreshToken          string
		RefreshTokenExpiresAt time.Time
		UserID                uuid.UUID
	}
	renewTokenImpl struct {
		tokenIssuer auth.TokenIssuer
		userRepo    repository.User
	}
)

func NewRenewToken(
	tokenIssuer auth.TokenIssuer,
	userRepo repository.User,
) RenewToken {
	return &renewTokenImpl{
		tokenIssuer: tokenIssuer,
		userRepo:    userRepo,
	}
}

func (u *renewTokenImpl) Execute(ctx context.Context, input RenewTokenInput) (*RenewTokenOutput, error) {
	// search user
	user, err := u.userRepo.GetByEmail(
		ctx,
		input.Email,
	)
	if err != nil {
		return nil, err
	}

	if err := password.Equals(
		user.Password().String(),
		input.RawPassword,
	); err != nil {
		return nil, err
	}
	// issue token
	accessToken, accessPayload, err := u.tokenIssuer.Create(
		user.ID(),
		config.Get().Token.AccessTokenDuration,
	)
	if err != nil {
		return nil, err
	}

	refreshToken, refreshPayload, err := u.tokenIssuer.Create(
		user.ID(),
		config.Get().Token.RefreshTokenDuration,
	)
	if err != nil {
		return nil, err
	}

	return &RenewTokenOutput{
		AccessToken:           accessToken,
		AccessTokenExpiresAt:  accessPayload.ExpiresAt(),
		RefreshToken:          refreshToken,
		RefreshTokenExpiresAt: refreshPayload.ExpiresAt(),
		UserID:                user.ID(),
	}, nil
}
