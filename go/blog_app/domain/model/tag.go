package model

import "fmt"

type Tag interface {
	ID() uint64
	Name() string
}

func ValidateTag(t Tag) error {
	if t.Name() == "" {
		return fmt.Errorf(
			"%w: tag must have a name",
			ErrConstructor,
		)
	}
	return nil
}
