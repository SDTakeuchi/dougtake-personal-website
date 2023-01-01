package model

import (
	"blog_app/adapter/persistance/database/postgres"
	"blog_app/domain/model"
	"blog_app/domain/model/uuid"
	"time"
)

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

func PostFromRecord(r postgres.Post) model.Post {
	return &post{
		id:        r.ID,
		title:     r.Title,
		body:      r.Body,
		userID:    r.UserID,
		tagIDs:    r.TagIDs,
		createdAt: r.CreatedAt,
		updatedAt: r.UpdatedAt,
	}
}

func PostToRecord(m model.Post) postgres.Post {
	return postgres.Post{
		ID:        m.ID(),
		Title:     m.Title(),
		Body:      m.Body(),
		UserID:    m.UserID(),
		TagIDs:    m.TagIDs(),
		CreatedAt: m.CreatedAt(),
		UpdatedAt: m.CreatedAt(),
	}
}