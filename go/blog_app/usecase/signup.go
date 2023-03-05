package usecase

import (
	modelimpl "blog_app/adapter/domain_impl/model"
	"blog_app/domain/model/password"
	"blog_app/domain/model/uuid"
	"blog_app/domain/repository"
	"context"
)

type (
	Signup interface {
		Execute(ctx context.Context, input SignupInput) (*SignupOutput, error)
	}
	SignupInput struct {
		Name        string
		Email       string
		RawPassword string
	}
	SignupOutput struct {
		ID    uuid.UUID
		Name  string
		Email string
	}
	signupImpl struct {
		userRepo repository.User
	}
)

func NewSignup(
	userRepo repository.User,
) Signup {
	return &signupImpl{
		userRepo: userRepo,
	}
}

func (u *signupImpl) Execute(ctx context.Context, input SignupInput) (*SignupOutput, error) {
	password, err := password.NewPassword(input.RawPassword)
	if err != nil {
		return nil, err
	}
	userInput, err := modelimpl.NewUser(
		input.Name,
		input.Email,
		*password,
	)
	if err != nil {
		return nil, err
	}
	user, err := u.userRepo.Create(
		ctx,
		userInput,
	)
	if err != nil {
		return nil, err
	}

	return &SignupOutput{
		ID:    user.ID(),
		Name:  user.Name(),
		Email: user.Email(),
	}, nil
}
