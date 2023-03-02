package usecase

import (
	modelimpl "blog_app/adapter/domain_impl/model"
	"blog_app/domain/model"
	"blog_app/domain/repository"
	"context"
	"time"
)

type (
	UpdateComment interface {
		Execute(ctx context.Context, input UpdateCommentInput) (*UpdateCommentOutput, error)
	}
	UpdateCommentInput struct {
		ID   uint64
		Body string
	}
	UpdateCommentOutput struct {
		Comment model.Comment
	}
	updateCommentImpl struct {
		commentRepo repository.Comment
	}
)

func NewUpdateComment(
	commentRepo repository.Comment,
) UpdateComment {
	return &updateCommentImpl{
		commentRepo: commentRepo,
	}
}

func (u *updateCommentImpl) Execute(ctx context.Context, input UpdateCommentInput) (*UpdateCommentOutput, error) {
	comment, err := u.commentRepo.Get(ctx, input.ID)
	if err != nil {
		return nil, err
	}
	newComment, err := modelimpl.NewComment(
		comment.ID(),
		input.Body,
		comment.PostID(),
		comment.CreatedAt(),
		time.Now(),
	)
	if err != nil {
		return nil, err
	}

	updatedComment, err := u.commentRepo.Update(ctx, newComment)
	if err != nil {
		return nil, err
	}
	return &UpdateCommentOutput{updatedComment}, nil
}
