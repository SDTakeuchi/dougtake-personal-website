package auth

import (
	"blog_app/domain/model/uuid"
	"time"
)

// an interface for managing tokens
type TokenIssuer interface {
	// creates a new token for a specific usernmame and duration
	// *Payload is neccessary to refresh token (cuz it needs token ID)
	Create(userID uuid.UUID, duration time.Duration) (string, Payload, error)

	// verifies if the token is valid or not
	Verify(token string) (Payload, error)
}
