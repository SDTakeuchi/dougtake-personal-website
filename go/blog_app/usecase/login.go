package usecase

import (
	"blog_app/adapter/config"
	authmodel "blog_app/adapter/domain_impl/model/auth"
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
		ClientIP    string
		UserAgent   string
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
		sessionRepo repository.Session
	}
)

func NewLogin(
	tokenIssuer auth.TokenIssuer,
	userRepo repository.User,
	sessionRepo repository.Session,
) Login {
	return &loginImpl{
		tokenIssuer: tokenIssuer,
		userRepo:    userRepo,
		sessionRepo: sessionRepo,
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
		auth.AccessToken,
		config.Get().Token.AccessTokenDuration,
	)
	if err != nil {
		return nil, err
	}

	refreshToken, refreshPayload, err := u.tokenIssuer.Create(
		user.ID(),
		auth.RefreshToken,
		config.Get().Token.RefreshTokenDuration,
	)
	if err != nil {
		return nil, err
	}

	err = createSession(
		ctx,
		u.sessionRepo,
		refreshToken,
		refreshPayload,
		input.ClientIP,
		input.UserAgent,
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

func createSession(
	ctx context.Context,
	sessionRepo repository.Session,
	refreshToken string,
	payload auth.Payload,
	clientIP string,
	userAgent string,
) error {
	// token is refresh token
	session, err := authmodel.NewSession(
		payload.ID(),
		payload.UserID(),
		refreshToken,
		userAgent,
		clientIP,
		payload.ExpiresAt(),
		payload.IssuedAt(),
	)
	if err != nil {
		return err
	}
	_, err = sessionRepo.Create(ctx, session)
	if err != nil {
		return err
	}
	return nil
}
