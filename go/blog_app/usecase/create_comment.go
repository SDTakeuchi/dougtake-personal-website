package usecase

import (
	"blog_app/domain/model"
	"blog_app/domain/repository"
	"context"
	"time"
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
		Comment model.Comment
	}
	createCommentImpl struct {
		postRepo    repository.Post
		commentRepo repository.Comment
	}
)

func NewCreateComment(
	postRepo repository.Post,
	commentRepo repository.Comment,
) CreateComment {
	return &createCommentImpl{
		postRepo:    postRepo,
		commentRepo: commentRepo,
	}
}

func (u *createCommentImpl) Execute(ctx context.Context, input CreateCommentInput) (*CreateCommentOutput, error) {
	post, err := u.postRepo.Get(ctx, input.PostID)
	if err != nil {
		return nil, err
	}

	now := time.Now()
	comment, err := model.NewComment(
		0,
		input.Body,
		post.ID(),
		now,
		now,
	)
	if err != nil {
		return nil, err
	}

	comment, err = u.commentRepo.Create(ctx, comment)
	if err != nil {
		return nil, err
	}
	return &CreateCommentOutput{comment}, nil
}
