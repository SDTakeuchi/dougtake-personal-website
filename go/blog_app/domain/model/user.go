package model

import "blog_app/domain/model/uuid"

type User interface {
	ID() uuid.UUID
	Name() string
	Email() string
	Password() string
}
