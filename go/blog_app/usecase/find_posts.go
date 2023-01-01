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
		Posts []postOutput
	}
	findPostsImpl struct {
		postRepo    repository.Post
		tagRepo     repository.Tag
		commentRepo repository.Comment
	}
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
		mPosts      []model.Post
		postOutputs []postOutput
	)

	// get posts
	if input.ID != 0 {
		post, err := u.postRepo.Get(ctx, input.ID)
		if err != nil {
			return nil, err
		}
		mPosts = append(mPosts, post)
	} else {
		posts, err := u.postRepo.Find(ctx, input.SearchChar, input.Offset, input.Limit)
		if err != nil {
			return nil, err
		}
		for _, p := range posts {
			mPosts = append(mPosts, p)
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
				tagOutput{t.ID(), t.Name()},
			)
		}

		//get comments
		mComment, err := u.commentRepo.FindByPostID(ctx, p.ID())
		if err != nil {
			return nil, err
		}

		var comments []commentOutput
		for _, c := range mComment {
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

	return &FindPostsOutput{postOutputs}, nil
}
