package usecase

import (
	"blog_app/domain/repository"
	"context"
)

type (
	DeleteComment interface {
		Execute(ctx context.Context, input DeleteCommentInput) (*DeleteCommentOutput, error)
	}
	DeleteCommentInput struct {
		ID uint64
	}
	DeleteCommentOutput struct {
	}
	deleteCommentImpl struct {
		commentRepo repository.Comment
	}
)

func NewDeleteComment(
	commentRepo repository.Comment,
) DeleteComment {
	return &deleteCommentImpl{
		commentRepo: commentRepo,
	}
}

func (u *deleteCommentImpl) Execute(ctx context.Context, input DeleteCommentInput) (*DeleteCommentOutput, error) {
	err := u.commentRepo.Delete(ctx, input.ID)
	if err != nil {
		return nil, err
	}
	return &DeleteCommentOutput{}, nil
}
