package repository

import (
	"blog_app/domain/model"
	"blog_app/domain/repository"
	modelimpl "blog_app/adapter/domain_impl/model"
	"blog_app/adapter/persistance/database/postgres"
	"context"
	"github.com/jinzhu/gorm"
)

type postRepository struct {
	db *gorm.DB
}

func NewPostRepository(db *gorm.DB) repository.Post {
	return postRepository{db}
}

const PostsMaxResponseSize = 10

func (r postRepository) Create(ctx context.Context, post model.Post) (model.Post, error) {
	return nil, nil
}

func (r postRepository) Get(ctx context.Context, id uint64) (model.Post, error) {
	var post postgres.Post
	if err := r.db.Take(&post, id).Error; err != nil {
		return nil, err
	}
	return modelimpl.PostFromRecord(post), nil
}

func (r postRepository) Find(
	ctx context.Context,
	tagID uint64,
	searchChar string,
	offset uint64,
	limit uint64,
) ([]model.Post, error) {
	var posts []postgres.Post

	if limit > PostsMaxResponseSize {
		limit = PostsMaxResponseSize
	}

    q := r.db.Order("created_at DESC")

	if tagID != 0 {
		q = q.Where("tag_id = ?", tagID)
	}

	if searchChar!= "" {
		q = q.Where("title LIKE ?", "%"+searchChar+"%")
	}

	if err := q.Offset(offset).Limit(limit).Find(&posts).Error; err != nil  {
		if err == gorm.ErrRecordNotFound {
			return nil, nil
		}
		return nil, err
	}

	var mPosts []model.Post
	for _, p := range posts {
		mPosts = append(mPosts, modelimpl.PostFromRecord(p))
	}

	return mPosts, nil
}

func (r postRepository) Update(ctx context.Context, post model.Post) (model.Post, error) {
	return nil, nil
}

func (r postRepository) Delete(ctx context.Context, id uint64) error {
	return nil
}
