package registry

import (
	"blog_app/adapter/handler"
	"blog_app/adapter/persistance/database/postgres"
)

type Registry struct {
	DBConn      postgres.DB
	PostHandler handler.PostHandler
}

func NewRegistry(db postgres.DB, postHandler handler.PostHandler) *Registry {
	return &Registry{
		db,
		postHandler,
	}
}
