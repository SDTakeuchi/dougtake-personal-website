//go:generate mockgen -source=session.go -destination=mock/session.go
package repository

import (
	"blog_app/domain/model/auth"
	"blog_app/domain/model/uuid"
	"context"
)

type Session interface {
	Create(ctx context.Context, session auth.Session) (auth.Session, error)
	Get(ctx context.Context, id uuid.UUID) (auth.Session, error)
}
