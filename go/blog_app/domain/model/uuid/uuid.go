package uuid

import (
	"context"
	"github.com/google/uuid"
)

type UUID struct {
	uuid.UUID
}

func NewV4(ctx context.Context) (UUID, error) {
	uid, err := uuid.NewV4()
	if err != nil {
		log.Errorf(ctx, "failed to generate UUID: %v", err)
		return UUID{}, err
	}
	return UUID{uid}, nil
}

func FromString(uuidString string) (UUID, error) {
	u, err := uuid.FromString(uuidString)
	if err != nil {
		return UUID{u}, err
	}
	return UUID{u}, nil
}
