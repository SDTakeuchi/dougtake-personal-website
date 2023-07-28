//go:generate mockgen -source=user.go -destination=mock/user.go
package repository

import (
	"blog_app/domain/model"
	"blog_app/domain/model/uuid"
	"context"
)

type User interface {
	Create(ctx context.Context, user model.User) (model.User, error)
	GetByEmail(ctx context.Context, email string) (model.User, error)
	GetByID(ctx context.Context, id uuid.UUID) (model.User, error)
	// Update(ctx context.Context, user model.User) (model.User, error)
	// Delete(ctx context.Context, id uuid.UUID) error
}
