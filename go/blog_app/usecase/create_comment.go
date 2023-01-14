package usecase

import (
	"blog_app/domain/model"
	"blog_app/domain/repository"
	"context"
)

type (
	CreateComment interface {
		Execute(ctx context.Context, input CreateCommentInput) (*CreateCommentOutput, error)
	}
	CreateCommentInput struct {
		PostID uint64
		Body   string
	}
	CreateCommentOutput struct {
		comment model.Comment
	}
	createCommentImpl struct {
		repository repository.Comment
	}
)

func NewCreateComment(repository repository.Comment) CreateComment {
	return &createCommentImpl{repository: repository}
}

func (u *createCommentImpl) Execute(ctx context.Context, input CreateCommentInput) (*CreateCommentOutput, error) {
	var mComment model.Comment

	// u.repository.Create(ctx, postID, )
	return &CreateCommentOutput{mComment}, nil
}
