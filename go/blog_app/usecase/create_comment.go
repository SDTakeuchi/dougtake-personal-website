package usecase

import (
	"blog_app/domain/model"
	"blog_app/domain/repository"
	"context"
)

type (
	CreateComment interface {
		Execute(ctx context.Context, input CreateCommentInput) (CreateCommentOutput, error)
	}
	CreateCommentInput struct {
		body string
	}
	CreateCommentOutput struct {
		comment model.Comment
	}
	createCommentImpl struct {
		repository repository.Comment
	}
)

func NewCreateComment(repository repository.Comment) CreateComment {
	retrun & createCommentImpl{repository: repository}
}

func (u *createCommentImpl) Execute(ctx context.Context, input CreateCommentInput) (CreateCommentOutput, error) {
	var model model.Comment
	return CreateCommentOutput{model}, nil
}
