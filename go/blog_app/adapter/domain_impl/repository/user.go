package repository

import (
	modelimpl "blog_app/adapter/domain_impl/model"
	"blog_app/adapter/persistance/database/postgres"
	"blog_app/domain/model"
	"blog_app/domain/repository"
	"context"
)

type userRepository struct {
	db *postgres.DB
}

func NewUserRepository(db *postgres.DB) repository.User {
	return &userRepository{db: db}
}

func (r *userRepository) Create(ctx context.Context, user model.User) (model.User, error) {
	pUser := modelimpl.UserToRecord(user)
	if err := r.db.Conn.WithContext(ctx).Create(&pUser).Error; err != nil {
		return nil, err
	}
	return modelimpl.UserFromRecord(pUser), nil
}

func (r *userRepository) GetByEmail(ctx context.Context, email string) (model.User, error) {
	var pUser postgres.User
	if err := r.db.Conn.WithContext(ctx).Where("email = ?", email).Take(&pUser).Error; err != nil {
		return nil, err
	}
	return modelimpl.UserFromRecord(pUser), nil
}
