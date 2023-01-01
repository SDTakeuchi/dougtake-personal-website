package registry

import "blog_app/adapter/persistance/database/postgres"

type Registry struct {
	DBConn postgres.DB
}
