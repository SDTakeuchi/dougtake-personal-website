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
	pComment := modelimpl.CommentToRecord(comment)
	if err := r.db.Conn.WithContext(ctx).Create(&pComment).Error; err != nil {
		return nil, err
	}
	return modelimpl.CommentFromRecord(pComment), nil
}

func (r *commentRepository) FindByPostID(ctx context.Context, postID uint64) ([]model.Comment, error) {
	var (
		pComments []postgres.Comment
		mComments []model.Comment
	)

	if err := r.db.Conn.WithContext(ctx).Order("created_at").Where("post_id = ?", postID).Find(&pComments).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return mComments, nil
		}
		return nil, err
	}

	for _, c := range pComments {
		mComments = append(mComments, modelimpl.CommentFromRecord(c))
	}
	return mComments, nil
}

func (r *commentRepository) Update(ctx context.Context, comment model.Comment) (model.Comment, error) {
	pComment := modelimpl.CommentToRecord(comment)
	if err := r.db.Conn.WithContext(ctx).Save(&pComment).Error; err != nil {
		return nil, err
	}
	return modelimpl.CommentFromRecord(pComment), nil
}

func (r *commentRepository) Delete(ctx context.Context, id uint64) error {
	return r.db.Conn.WithContext(ctx).Delete(&postgres.Comment{}, id).Error
}
