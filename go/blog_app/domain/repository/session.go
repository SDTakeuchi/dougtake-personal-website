package repository

import (
	"blog_app/domain/model/auth"
	"context"
)

type Session interface {
	Get(ctx context.Context, id string) (auth.Session, error)
	Create(ctx context.Context, session auth.Session) (auth.Session, error)
}
