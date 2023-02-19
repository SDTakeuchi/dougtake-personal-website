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

func NewTag(id uint64, name string) (model.Tag, error) {
	t := &tag{
		id:   id,
		name: name,
	}
	if err := model.ValidateTag(t); err != nil {
		return nil, err
	}
	return t, nil
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
