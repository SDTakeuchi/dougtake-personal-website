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
	Login interface {
		Execute(ctx context.Context, input LoginInput) (*LoginOutput, error)
	}
	LoginInput struct {
		Email       string
		RawPassword string
	}
	LoginOutput struct {
		AccessToken           string
		AccessTokenExpiresAt  time.Time
		RefreshToken          string
		RefreshTokenExpiresAt time.Time
		UserID                uuid.UUID
	}
	loginImpl struct {
		tokenIssuer auth.TokenIssuer
		userRepo    repository.User
	}
)

func NewLogin(
	tokenIssuer auth.TokenIssuer,
	userRepo repository.User,
) Login {
	return &loginImpl{
		tokenIssuer: tokenIssuer,
		userRepo:    userRepo,
	}
}

func (u *loginImpl) Execute(ctx context.Context, input LoginInput) (*LoginOutput, error) {
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

	return &LoginOutput{
		AccessToken:           accessToken,
		AccessTokenExpiresAt:  accessPayload.ExpiresAt(),
		RefreshToken:          refreshToken,
		RefreshTokenExpiresAt: refreshPayload.ExpiresAt(),
		UserID:                user.ID(),
	}, nil
}
