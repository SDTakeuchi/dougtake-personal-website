package model

import (
	"fmt"
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
		return fmt.Errorf(
			"%w: comment body must not be empty",
			ErrConstructor,
		)
	}
	if c.PostID() == 0 {
		return fmt.Errorf(
			"%w: comment postID must not be empty",
			ErrConstructor,
		)
	}
	if c.CreatedAt().IsZero() && !c.UpdatedAt().IsZero() {
		return fmt.Errorf(
			"%w: comment createdAt must not be empty when updatedAt is filled",
			ErrConstructor,
		)
	}
	return nil
}
