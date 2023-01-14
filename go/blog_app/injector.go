//go:build wireinject

package main

import "github.com/google/wire"

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
		NewPostHandler,
	)
	return Event{}
}
