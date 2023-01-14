//go:build wireinject
// +build wireinject

package main

import (
	"blog_app/adapter/domain_impl/repository"
	"blog_app/adapter/handler"
	"blog_app/adapter/persistance/database/postgres"
	"blog_app/adapter/registry"
	"blog_app/usecase"

	"github.com/google/wire"
)

func initialize(db *postgres.DB) *registry.Registry {
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

		registry.NewRegistry,
	)

	return &registry.Registry{}
}
