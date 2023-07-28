package model

import (
	"blog_app/adapter/persistance/database/postgres"
	"blog_app/domain/model"
	"blog_app/domain/model/uuid"
	"time"

	"github.com/lib/pq"
)

type post struct {
	id        uint64
	title     string
	body      string
	userID    uuid.UUID
	tagIDs    []uint64
	createdAt time.Time `gorm:"autoCreateTime"`
	updatedAt time.Time `gorm:"autoCreateTime"`
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
) (model.Post, error) {
	p := &post{
		id:        id,
		title:     title,
		body:      body,
		userID:    userID,
		tagIDs:    tagIDs,
		createdAt: createdAt,
		updatedAt: updatedAt,
	}
	if err := model.ValidatePost(p); err != nil {
		return nil, err
	}
	return p, nil
}

func PostFromRecord(r postgres.Post) model.Post {
	tagIDsUint64 := make([]uint64, len(r.TagIDs))
    for i, v := range r.TagIDs {
        tagIDsUint64[i] = uint64(v)
    }
	return &post{
		id:        r.ID,
		title:     r.Title,
		body:      r.Body,
		userID:    r.UserID,
		tagIDs:    tagIDsUint64,
		createdAt: r.CreatedAt,
		updatedAt: r.UpdatedAt,
	}
}

func PostToRecord(m model.Post) postgres.Post {
	tagIDsInt64 := make([]int64, len(m.TagIDs()))
    for i, v := range m.TagIDs() {
        tagIDsInt64[i] = int64(v)
    }
	return postgres.Post{
		ID:        m.ID(),
		Title:     m.Title(),
		Body:      m.Body(),
		UserID:    m.UserID(),
		TagIDs:    pq.Int64Array(tagIDsInt64),
		CreatedAt: m.CreatedAt(),
		UpdatedAt: m.CreatedAt(),
	}
}
