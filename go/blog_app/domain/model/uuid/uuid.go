package uuid

import (
	"github.com/google/uuid"
)

type UUID struct {
	uuid.UUID
}

func New() UUID {
	id := uuid.New()
	return UUID{id}
}

func FromString(uuidString string) (UUID, error) {
	u, err := uuid.Parse(uuidString)
	return UUID{u}, err
}
