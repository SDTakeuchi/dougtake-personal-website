package model

import (
	"blog_app/adapter/persistance/database/postgres"
	"blog_app/domain/model"
)

func PostFromRecord(r postgres.Post) (model.Post, error) {
	return model.NewPost(
		r.ID,
		r.Title,
		r.Body,
		r.UserID,
		r.TagIDs,
		r.CreatedAt,
		r.UpdatedAt,
	)
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
