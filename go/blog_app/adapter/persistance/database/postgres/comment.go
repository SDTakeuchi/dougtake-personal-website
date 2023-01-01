package postgres

import "time"

type Comment struct {
	ID        uint64
	Body      string
	PostID    uint64
	CreatedAt time.Time
	UpdatedAt time.Time
}
