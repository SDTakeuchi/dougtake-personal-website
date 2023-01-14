package postgres

import "blog_app/domain/model/uuid"

type User struct {
	ID       uuid.UUID
	Name     string
	Email    string
	Password string
}
