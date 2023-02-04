package repository

import (
	modelimpl "blog_app/adapter/domain_impl/model"
	"blog_app/adapter/persistance/database/postgres"
	"blog_app/domain/model"
	"blog_app/domain/repository"
	"context"
	"time"

	"gorm.io/gorm"
)

type postRepository struct {
	db *postgres.DB
}

func NewPostRepository(db *postgres.DB) repository.Post {
	return &postRepository{db}
}

func (r *postRepository) Create(ctx context.Context, post model.Post) (model.Post, error) {
	pPost := modelimpl.PostToRecord(post)
	now := time.Now()
	pPost.CreatedAt = now
	pPost.UpdatedAt = now
	if err := r.db.Conn.WithContext(ctx).Create(&pPost).Error; err != nil {
		return nil, err
	}
	return modelimpl.PostFromRecord(pPost), nil
}

func (r *postRepository) Get(ctx context.Context, id uint64) (model.Post, error) {
	var post postgres.Post
	if err := r.db.Conn.WithContext(ctx).Take(&post, id).Error; err != nil {
		return nil, err
	}
	return modelimpl.PostFromRecord(post), nil
}

func (r *postRepository) Find(
	ctx context.Context,
	tagID uint64,
	searchChar string,
	offset uint64,
	limit uint64,
) ([]model.Post, error) {
	q := r.db.Conn.WithContext(ctx).Order("created_at DESC")

	if tagID != 0 {
		q = q.Where("tag_id = ?", tagID)
	}

	if searchChar != "" {
		param := "%" + searchChar + "%"
		q = q.Where("body LIKE ?", param)
	}

	var pPosts []postgres.Post
	if err := q.Offset(int(offset)).Limit(int(limit)).Find(&pPosts).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, nil
		}
		return nil, err
	}

	var mPosts []model.Post
	for _, p := range pPosts {
		mPosts = append(mPosts, modelimpl.PostFromRecord(p))
	}

	return mPosts, nil
}

func (r *postRepository) Update(ctx context.Context, post model.Post) (model.Post, error) {
	pPost := modelimpl.PostToRecord(post)
	// pPost.UpdatedAt gets updated by defalt by Gorm
	if err := r.db.Conn.WithContext(ctx).Save(&pPost).Error; err != nil {
		return nil, err
	}
	return modelimpl.PostFromRecord(pPost), nil
}

func (r *postRepository) Delete(ctx context.Context, id uint64) error {
	return r.db.Conn.WithContext(ctx).Delete(&postgres.Post{}, id).Error
}
