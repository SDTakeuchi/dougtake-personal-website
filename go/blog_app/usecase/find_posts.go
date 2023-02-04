package usecase

import (
	"blog_app/domain/model"
	"blog_app/domain/repository"
	"context"
	"time"
)

type (
	FindPosts interface {
		Execute(ctx context.Context, input FindPostsInput) (*FindPostsOutput, error)
	}
	FindPostsInput struct {
		ID         uint64
		TagID      uint64
		SearchChar string
		Offset     uint64
		Limit      uint64
	}
	tagOutput struct {
		ID   uint64
		Name string
	}
	commentOutput struct {
		ID        uint64
		Body      string
		CreatedAt time.Time
		UpdatedAt time.Time
	}
	postOutput struct {
		ID        uint64
		Title     string
		Body      string
		CreatedAt time.Time
		UpdatedAt time.Time
		Tags      []tagOutput
		Comments  []commentOutput
	}
	FindPostsOutput struct {
		Posts         []postOutput
		NextPostIndex uint64
	}
	findPostsImpl struct {
		postRepo    repository.Post
		tagRepo     repository.Tag
		commentRepo repository.Comment
	}
)

const (
	defaultResponseSize = 10
	maxResponseSize     = 20
)

func NewFindPosts(
	postRepo repository.Post,
	tagRepo repository.Tag,
	commentRepo repository.Comment,
) FindPosts {
	return &findPostsImpl{
		postRepo:    postRepo,
		tagRepo:     tagRepo,
		commentRepo: commentRepo,
	}
}

func (u *findPostsImpl) Execute(ctx context.Context, input FindPostsInput) (*FindPostsOutput, error) {
	var (
		mPosts        []model.Post
		postOutputs   []postOutput
		nextPostIndex uint64
	)

	if input.ID != 0 {
		post, err := u.postRepo.Get(ctx, input.ID)
		if err != nil {
			return nil, err
		}
		mPosts = append(mPosts, post)
	} else {
		limit := input.Limit
		if limit == 0 {
			limit = defaultResponseSize
		}
		if limit > maxResponseSize {
			limit = maxResponseSize
		}

		// increment limit to check if there is at least one more post
		posts, err := u.postRepo.Find(ctx, input.TagID, input.SearchChar, input.Offset, limit+1)
		if err != nil {
			return nil, err
		}
		mPosts = append(mPosts, posts...)

		// for pagination
		if len(mPosts) > int(limit) {
			nextPostIndex = mPosts[len(mPosts)-1].ID()
			mPosts = mPosts[:len(mPosts)-1]
		}
	}

	for _, p := range mPosts {
		// get tags
		mTags, err := u.tagRepo.Find(ctx, p.TagIDs())
		if err != nil {
			return nil, err
		}

		var tags []tagOutput
		for _, t := range mTags {
			tags = append(
				tags,
				tagOutput{
					t.ID(),
					t.Name(),
				},
			)
		}

		//get comments
		mComments, err := u.commentRepo.FindByPostID(ctx, p.ID())
		if err != nil {
			return nil, err
		}

		var comments []commentOutput
		for _, c := range mComments {
			comments = append(
				comments,
				commentOutput{
					c.ID(),
					c.Body(),
					c.CreatedAt(),
					c.UpdatedAt(),
				},
			)
		}

		// append to results
		postOutputs = append(
			postOutputs,
			postOutput{
				ID:        p.ID(),
				Title:     p.Title(),
				Body:      p.Body(),
				CreatedAt: p.CreatedAt(),
				UpdatedAt: p.UpdatedAt(),
				Tags:      tags,
				Comments:  comments,
			},
		)
	}

	return &FindPostsOutput{
		postOutputs,
		nextPostIndex,
	}, nil
}
