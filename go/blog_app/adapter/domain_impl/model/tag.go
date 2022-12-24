package model

import (
	"blog_app/adapter/persistance/database/postgres"
	"blog_app/domain/model"
)

type tag struct {
	id   uint64
	name string
}

func (t *tag) ID() uint64 { return t.id }

func (t *tag) Name() string { return t.name }

func NewTag(id uint64, name string) model.Tag {
	return &tag{id, name}
}

func TagFromRecord(r postgres.Tag) model.Tag {
	return &tag{
		id:   r.ID,
		name: r.Name,
	}
}

func TagToRecord(m model.Tag) postgres.Tag {
	return postgres.Tag{
		ID:   m.ID(),
		Name: m.Name(),
	}
}
