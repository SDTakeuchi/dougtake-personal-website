//go:generate mockgen -source=post.go -destination=../../adapter/domain_impl/repository/mock/post.go
package repository

import (
	"blog_app/domain/model"
	"context"
)

type Post interface {
	Create(ctx context.Context, post model.Post) (model.Post, error)
	Get(ctx context.Context, id uint64) (model.Post, error)
	Find(ctx context.Context, tagID uint64, searchChar string, offset, limit uint64) ([]model.Post, error)
	Update(ctx context.Context, post model.Post) (model.Post, error)
	Delete(ctx context.Context, id uint64) error
}
