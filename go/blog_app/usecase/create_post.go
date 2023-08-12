package usecase

import (
	modelimpl "blog_app/adapter/domain_impl/model"
	"blog_app/domain/model"
	"blog_app/domain/model/uuid"
	"blog_app/domain/repository"
	"context"
	"time"
)

type (
	CreatePost interface {
		Execute(ctx context.Context, input CreatePostInput) (*CreatePostOutput, error)
	}
	CreatePostInput struct {
		UserID string
		Title  string
		Body   string
		TagIDs []uint64
	}
	CreatePostOutput struct {
		Post model.Post
	}
	createPostImpl struct {
		userRepo repository.User
		tagRepo  repository.Tag
		postRepo repository.Post
	}
)

func NewCreatePost(
	userRepo repository.User,
	tagRepo repository.Tag,
	postRepo repository.Post,
) CreatePost {
	return &createPostImpl{
		userRepo: userRepo,
		tagRepo:  tagRepo,
		postRepo: postRepo,
	}
}

func (u *createPostImpl) Execute(ctx context.Context, input CreatePostInput) (*CreatePostOutput, error) {
	userID, err := uuid.Parse(input.UserID)
	if err != nil {
		return nil, err
	}
	if _, err := u.userRepo.GetByID(ctx, userID); err != nil {
		return nil, err
	}

	if _, err = u.tagRepo.Find(ctx, input.TagIDs); err != nil {
		return nil, err
	}

	post, err := modelimpl.NewPost(
		0,
		input.Title,
		input.Body,
		userID,
		input.TagIDs,
		time.Now(),
		time.Time{},
	)
	if err != nil {
		return nil, err
	}

	post, err = u.postRepo.Create(ctx, post)
	if err != nil {
		return nil, err
	}
	return &CreatePostOutput{Post: post}, nil
}
