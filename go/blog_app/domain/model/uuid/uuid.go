package uuid

import (
	"github.com/google/uuid"
)

type UUID struct {
	uuid.UUID
}

func (u *UUID) String() string {
	return u.UUID.String()
}

func New() UUID {
	id := uuid.New()
	return UUID{id}
}

func Parse(uuidString string) (UUID, error) {
	u, err := uuid.Parse(uuidString)
	return UUID{u}, err
}
