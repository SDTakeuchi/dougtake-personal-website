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

func ValidatePost(p Post) error {
	if p.Title() == "" {
		return errors.New("post title must not be empty")
	}
	if p.Body() == "" {
		return errors.New("post body must not be empty")
	}
	if p.CreatedAt().IsZero() && !p.UpdatedAt().IsZero() {
		return errors.New("post createdAt must not be empty when updatedAt is filled")
	}
	return nil
}
