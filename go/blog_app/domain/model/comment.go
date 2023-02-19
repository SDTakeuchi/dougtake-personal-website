package model

import (
	"errors"
	"time"
)

type Comment interface {
	ID() uint64
	Body() string
	PostID() uint64
	CreatedAt() time.Time
	UpdatedAt() time.Time
}

func ValidateComment(c Comment) error {
	if c.Body() == "" {
		return errors.New("comment body must not be empty")
	}
	if c.PostID() == 0 {
		return errors.New("comment postID must not be empty")
	}
	if c.CreatedAt().IsZero() && !c.UpdatedAt().IsZero() {
		return errors.New("comment createdAt must not be empty when updatedAt is filled")
	}
	return nil
}
