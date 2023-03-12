package middleware

import (
	"context"
)

func (s *userService) AuthenticateUser(ctx context.Context) error {
	// substruct session from context
	accessToken, err := //TODO

	session, err := s.sessionRepo.Get(ctx, accessToken)
	if err != nil {
		return err
	}
	return nil
}
