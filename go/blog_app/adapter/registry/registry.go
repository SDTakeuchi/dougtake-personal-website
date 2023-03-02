package registry

import (
	"blog_app/adapter/handler"
	"blog_app/adapter/persistance/database/postgres"
)

type Registry struct {
	DBConn         *postgres.DB
	PostHandler    handler.PostHandler
	CommentHandler handler.CommentHandler
	TagHandler     handler.TagHandler
}

func NewRegistry(
	db *postgres.DB,
	postHandler handler.PostHandler,
	commentHandler handler.CommentHandler,
	tagHandler handler.TagHandler,
) Registry {
	return Registry{
		DBConn:         db,
		PostHandler:    postHandler,
		CommentHandler: commentHandler,
		TagHandler:     tagHandler,
	}
}
