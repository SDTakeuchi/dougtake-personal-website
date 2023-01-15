package repository

import (
	modelimpl "blog_app/adapter/domain_impl/model"
	"blog_app/adapter/persistance/database/postgres"
	"blog_app/domain/model"
	"blog_app/domain/repository"
	"context"

	"gorm.io/gorm"
)

type commentRepository struct {
	db *postgres.DB
}

func NewCommentRepository(db *postgres.DB) repository.Comment {
	return &commentRepository{db}
}

func (r *commentRepository) Create(ctx context.Context, comment model.Comment) (model.Comment, error) {
	return nil, nil
}

func (r *commentRepository) FindByPostID(ctx context.Context, postID uint64) ([]model.Comment, error) {
	var comments []postgres.Comment

	if err := r.db.Conn.Order("created_at").Find(&comments).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, nil
		}
		return nil, err
	}

	var mComments []model.Comment
	for _, c := range comments {
		mComments = append(mComments, modelimpl.CommentFromRecord(c))
	}

	return mComments, nil
}

func (r *commentRepository) Update(ctx context.Context, comment model.Comment) (model.Comment, error) {
	return nil, nil
}

func (r *commentRepository) Delete(ctx context.Context, id uint64) error {
	return nil
}
