package model

import (
	"blog_app/domain/model/uuid"
	"errors"
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

type post struct {
	id        uint64
	title     string
	body      string
	userID    uuid.UUID
	tagIDs    []uint64
	createdAt time.Time
	updatedAt time.Time
}

func (p *post) ID() uint64 { return p.id }

func (p *post) Title() string { return p.title }

func (p *post) Body() string { return p.body }

func (p *post) UserID() uuid.UUID { return p.userID }

func (p *post) TagIDs() []uint64 { return p.tagIDs }

func (p *post) CreatedAt() time.Time { return p.createdAt }

func (p *post) UpdatedAt() time.Time { return p.updatedAt }

func NewPost(
	id uint64,
	title, body string,
	userID uuid.UUID,
	tagIDs []uint64,
	createdAt, updatedAt time.Time,
) (Post, error) {
	if title == "" {
		return nil, errors.New("post title must not be empty")
	}
	if body == "" {
		return nil, errors.New("post body must not be empty")
	}
	if createdAt.IsZero() {
		if !updatedAt.IsZero() {
			return nil, errors.New("invalid data: createdAt must be filed when updatedAt is filled")
		}
		createdAt = time.Now()
		updatedAt = time.Now()
	}
	return &post{
		id:        id,
		title:     title,
		body:      body,
		userID:    userID,
		tagIDs:    tagIDs,
		createdAt: createdAt,
		updatedAt: updatedAt,
	}, nil
}
