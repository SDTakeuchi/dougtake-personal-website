package model

import (
	"blog_app/domain/model/password"
	"blog_app/domain/model/uuid"
	"fmt"
	"regexp"
)

type User interface {
	ID() uuid.UUID
	Name() string
	Email() string
	Password() password.Password
}

func ValidateUser(u User) error {
	if u.Name() == "" {
		return fmt.Errorf(
			"%w: user name is required",
			ErrConstructor,
		)
	}
	if !validateEmail(u.Email()) {
		return fmt.Errorf(
			"%w: invalid email",
			ErrConstructor,
		)
	}
	if u.Password().String() == "" {
		return fmt.Errorf(
			"%w: password is required",
			ErrConstructor,
		)
	}
	return nil
}

func validateEmail(email string) bool {
	pattern := `^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`
	regex := regexp.MustCompile(pattern)
	return regex.MatchString(email)
}
