package model

import (
	"errors"
	"time"
)

type Comment interface {
	ID() uint64
	Body() string
	PostID() uint64
	CreatedAt() time.Time
	UpdatedAt() time.Time
}

type comment struct {
	id        uint64
	body      string
	postID    uint64
	createdAt time.Time
	updatedAt time.Time
}

func (c *comment) ID() uint64 { return c.id }

func (c *comment) Body() string { return c.body }

func (c *comment) PostID() uint64 { return c.id }

func (c *comment) CreatedAt() time.Time { return c.createdAt }

func (c *comment) UpdatedAt() time.Time { return c.updatedAt }

func NewComment(
	id uint64,
	body string,
	postID uint64,
	createdAt, updatedAt time.Time,
) (Comment, error) {
	if body == "" {
		return nil, errors.New("comment body must not be empty")
	}
	if postID == 0 {
		return nil, errors.New("comment postID must not be empty")
	}
	return &comment{
		id:        id,
		body:      body,
		postID:    postID,
		createdAt: createdAt,
		updatedAt: updatedAt,
	}, nil
}
