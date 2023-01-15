//go:generate mockgen -source=comment.go -destination=../../adapter/domain_impl/repository/mock/comment.go
package repository

import (
	"blog_app/domain/model"
	"context"
)

type Comment interface {
	Create(ctx context.Context, comment model.Comment) (model.Comment, error)
	FindByPostID(ctx context.Context, postID uint64) ([]model.Comment, error)
	Update(ctx context.Context, comment model.Comment) (model.Comment, error)
	Delete(ctx context.Context, id uint64) error
}
