package usecase

import (
	"blog_app/domain/model/password"
	"blog_app/domain/repository"
	"context"
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
		sessionKey []byte
	}
	loginImpl struct {
		userRepo repository.User
	}
)

func NewLogin(
	userRepo repository.User,
) Login {
	return &loginImpl{
		userRepo: userRepo,
	}
}

func (u *loginImpl) Execute(ctx context.Context, input LoginInput) (*LoginOutput, error) {
	user, err := u.userRepo.GetByEmail(
		ctx,
		input.Email,
	)
	if err != nil {
		return nil, err
	}

	if err := password.Equals(
		user.HashedPassword(),
		input.RawPassword,
	); err != nil {
		return nil, err
	}

	return &LoginOutput{tags}, nil
}

