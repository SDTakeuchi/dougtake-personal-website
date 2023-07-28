//go:build wireinject
// +build wireinject

package main

import (
	"blog_app/adapter/domain_impl/repository"
	"blog_app/adapter/handler"
	"blog_app/adapter/persistance/database/postgres"
	"blog_app/adapter/registry"
	"blog_app/domain/model/auth"
	"blog_app/usecase"

	"github.com/google/wire"
)

func initialize(db *postgres.DB, tokenIssuer auth.TokenIssuer) registry.Registry {
	wire.Build(
		// user
		repository.NewUserRepository,
		// auth
		repository.NewSessionRepository,
		usecase.NewSignup,
		usecase.NewLogin,
		usecase.NewRenewToken,
		handler.NewAuthHandler,
		// tag
		repository.NewTagRepository,
		usecase.NewGetTags,
		handler.NewTagHandler,
		// comment
		repository.NewCommentRepository,
		usecase.NewCreateComment,
		usecase.NewUpdateComment,
		usecase.NewDeleteComment,
		handler.NewCommentHandler,
		// post
		repository.NewPostRepository,
		usecase.NewCreatePost,
		usecase.NewFindPosts,
		handler.NewPostHandler,
		//registry
		registry.NewRegistry,
	)

	return registry.Registry{}
}
