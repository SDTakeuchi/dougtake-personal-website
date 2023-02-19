package model

import "errors"

type Tag interface {
	ID() uint64
	Name() string
}

func ValidateTag(t Tag) error {
	if t.Name() == "" {
		return errors.New("tag must have a name")
	}
	return nil
}
