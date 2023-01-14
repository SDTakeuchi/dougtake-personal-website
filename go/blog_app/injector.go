//go:build wireinject

package main

import (
	"blog_app/adapter/handler"
	"blog_app/domain/repository"
	"blog_app/usecase"

	"github.com/google/wire"
)

func InitializeEvent() Event {
	wire.Build(
		// user
		// tag
		repository.NewTagRepository,
		// comment
		repository.NewCommentRepository,
		// post
		repository.NewPostRepository,
		usecase.NewFindPosts,
		handler.NewPostHandler,
	)
	return Event{}
}
