package service

import (
	"blog_app/usecase"
	"context"
)

type PostQueryService interface {
	FindPosts(ctx context.Context, searchChar string, tagID uint64, limit, offset uint64) (usecase.FindPostsOutput, error)
}
