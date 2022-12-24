package repository

import (
	"blog_app/domain/model"
	"context"
)

type Post interface {
	Create(ctx context.Context, post model.Post) (model.Post, error)
	Get(ctx context.Context, id uint64) (model.Post, error)
	Find(ctx context.Context, searchChar string, offset, limit uint64) ([]model.Post, error)
	Update(ctx context.Context, post model.Post) (model.Post, error)
	Delete(ctx context.Context, id uint64) error
}
