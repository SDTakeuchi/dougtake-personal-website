package registry

import (
	"blog_app/adapter/handler"
	"blog_app/adapter/persistance/database/postgres"
)

type Registry struct {
	DBConn      *postgres.DB
	PostHandler handler.PostHandler
	CommentHandler handler.CommentHandler
}

func NewRegistry(
	db *postgres.DB,
	postHandler handler.PostHandler,
	commentHandler handler.CommentHandler,
) Registry {
	return Registry{
		db,
		postHandler,
		commentHandler,
	}
}
