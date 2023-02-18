package model

import (
	"blog_app/adapter/persistance/database/postgres"
	"blog_app/domain/model"
)

func CommentFromRecord(r postgres.Comment) (model.Comment, error) {
	return model.NewComment(
		r.ID,
		r.Body,
		r.PostID,
		r.CreatedAt,
		r.UpdatedAt,
	)
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
