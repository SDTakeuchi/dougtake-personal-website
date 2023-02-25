package service

import (
	"blog_app/adapter/persistance/database/postgres"
	"blog_app/domain/service"
	"blog_app/usecase"
	"context"
	"time"
)

type (
	res struct {
		commentBody      string
		commentCreatedAt time.Time
		commentUpdatedAt time.Time
		postID           uint64
		postTitle        string
		postBody         string
		postCreatedAt    time.Time
		postUpdatedAt    time.Time
		tagName          string
	}
	findPostsService struct {
		db *postgres.DB
	}
)

func NewFindPostsService(db *postgres.DB) service.PostQueryService {
	return &findPostsService{db: db}
}

func (s *findPostsService) FindPosts(
	ctx context.Context,
	searchChar string,
	tagID uint64,
	limit, offset uint64,
) (usecase.FindPostsOutput, error) {
	var result res
	if err := s.db.Conn.WithContext(ctx).Raw(`
	SELECT
		c.body AS comment_body,
		c.created_at AS comment_created_at,
		c.updated_at AS comment_updated_at,
		p.id AS post_id,
		p.title AS post_title,
		p.body AS post_body,
		p.created_at AS post_created_at,
		p.updated_at AS post_updated_at,
		t.name AS tag_name
	FROM
		comments AS c
		INNER JOIN posts AS p ON c.post_id = p.id
		LEFT JOIN tags AS t ON p.tag_id = t.id
	WHERE
		body LIKE '%?%'
	LIMIT ?
	OFFSET ?;
	`, searchChar, limit, offset).Scan(&result).Error; err != nil {
		return nil, err
	}
}
