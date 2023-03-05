package registry

import (
	"blog_app/adapter/handler"
	"blog_app/adapter/persistance/database/postgres"
	"blog_app/domain/model/auth"
)

type Registry struct {
	DBConn         *postgres.DB
	TokenIssuer    auth.TokenIssuer
	AuthHandler    handler.AuthHandler
	PostHandler    handler.PostHandler
	CommentHandler handler.CommentHandler
	TagHandler     handler.TagHandler
}

func NewRegistry(
	db *postgres.DB,
	tokenIssuer auth.TokenIssuer,
	authHandler handler.AuthHandler,
	postHandler handler.PostHandler,
	commentHandler handler.CommentHandler,
	tagHandler handler.TagHandler,
) Registry {
	return Registry{
		DBConn:         db,
		TokenIssuer:    tokenIssuer,
		AuthHandler:    authHandler,
		PostHandler:    postHandler,
		CommentHandler: commentHandler,
		TagHandler:     tagHandler,
	}
}
