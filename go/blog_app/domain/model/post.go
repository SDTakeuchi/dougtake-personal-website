package model

import (
	"blog_app/domain/model/uuid"
	"fmt"
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
		return fmt.Errorf(
			"%w: post title must not be empty",
			ErrInvalidParams,
		)
	}
	if p.Body() == "" {
		return fmt.Errorf(
			"%w: post body must not be empty",
			ErrInvalidParams,
		)
	}
	if p.CreatedAt().IsZero() && !p.UpdatedAt().IsZero() {
		return fmt.Errorf(
			"%w: post createdAt must not be empty when updatedAt is filled",
			ErrInvalidParams,
		)
	}
	return nil
}
