package model

import (
	"blog_app/domain/model/uuid"
	"time"
)

type Post interface {
	ID() uint64
	Title() string
	Body() string
	UserID() uuid.UUID
	TagIDs() []uint64
	CreatedAt() time.Time
	UpdatedAt() time.Time
}
