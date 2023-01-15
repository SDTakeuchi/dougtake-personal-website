//go:generate mockgen -source=tag.go -destination=mock/tag.go
package repository

import (
	"blog_app/domain/model"
	"context"
)

type Tag interface {
	Create(ctx context.Context, tag model.Tag) (model.Tag, error)
	Find(ctx context.Context, ids []uint64) ([]model.Tag, error)
	Update(ctx context.Context, tag model.Tag) (model.Tag, error)
	Delete(ctx context.Context, id uint64) error
}
