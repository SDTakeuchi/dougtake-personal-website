package model

import "time"

type Comment interface {
	ID() uint64
	Body() string
	PostID() uint64
	CreatedAt() time.Time
	UpdatedAt() time.Time
}
