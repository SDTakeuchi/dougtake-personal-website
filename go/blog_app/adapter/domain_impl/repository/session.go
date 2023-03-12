package repository

import (
	modelimpl "blog_app/adapter/domain_impl/model/auth"
	"blog_app/adapter/persistance/database/postgres"
	"blog_app/domain/model/auth"
	"blog_app/domain/model/uuid"
	"blog_app/domain/repository"
	"context"
)

type sessionRepository struct {
	db *postgres.DB
}

func NewSessionRepository(db *postgres.DB) repository.Session {
	return &sessionRepository{db}
}

func (r *sessionRepository) Create(ctx context.Context, session auth.Session) (auth.Session, error) {
	pSession := modelimpl.SessionToRecord(session)
	if err := r.db.Conn.WithContext(ctx).Create(&pSession).Error; err != nil {
		return nil, err
	}
	return modelimpl.SessionFromRecord(pSession), nil
}

func (r *sessionRepository) Get(ctx context.Context, id uuid.UUID) (auth.Session, error) {
	var pSession postgres.Session
	if err := r.db.Conn.WithContext(ctx).Take(&pSession, id.String()).Error; err != nil {
		return nil, err
	}
	return modelimpl.SessionFromRecord(pSession), nil
}
