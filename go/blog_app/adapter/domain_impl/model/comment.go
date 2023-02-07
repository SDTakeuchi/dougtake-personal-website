package model

import (
	"blog_app/adapter/persistance/database/postgres"
	"blog_app/domain/model"
	"errors"
	"time"
)

type comment struct {
	id        uint64
	body      string
	postID    uint64
	createdAt time.Time `gorm:"autoCreateTime"`
	updatedAt time.Time `gorm:"autoCreateTime"`
}

func (c *comment) ID() uint64 { return c.id }

func (c *comment) Body() string { return c.body }

func (c *comment) PostID() uint64 { return c.postID }

func (c *comment) CreatedAt() time.Time { return c.createdAt }

func (c *comment) UpdatedAt() time.Time { return c.updatedAt }

func NewComment(
	id uint64,
	body string,
	postID uint64,
	createdAt, updatedAt time.Time,
) (model.Comment, error) {
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

func CommentFromRecord(r postgres.Comment) model.Comment {
	return &comment{
		id:        r.ID,
		body:      r.Body,
		postID:    r.PostID,
		createdAt: r.CreatedAt,
		updatedAt: r.UpdatedAt,
	}
}

func CommentToRecord(m model.Comment) postgres.Comment {
	return postgres.Comment{
		ID:        m.ID(),
		Body:      m.Body(),
		PostID:    m.PostID(),
		CreatedAt: m.CreatedAt(),
		UpdatedAt: m.CreatedAt(),
	}
}
