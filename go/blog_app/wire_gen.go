// Code generated by Wire. DO NOT EDIT.

//go:generate go run github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package main

import (
	"blog_app/adapter/domain_impl/repository"
	"blog_app/adapter/handler"
	"blog_app/adapter/persistance/database/postgres"
	"blog_app/adapter/registry"
	"blog_app/usecase"
)

// Injectors from injector.go:

func initialize(db *postgres.DB) registry.Registry {
	post := repository.NewPostRepository(db)
	tag := repository.NewTagRepository(db)
	comment := repository.NewCommentRepository(db)
	findPosts := usecase.NewFindPosts(post, tag, comment)
	postHandler := handler.NewPostHandler(findPosts)
	createComment := usecase.NewCreateComment(post, comment)
	commentHandler := handler.NewCommentHandler(createComment)
	registryRegistry := registry.NewRegistry(db, postHandler, commentHandler)
	return registryRegistry
}
