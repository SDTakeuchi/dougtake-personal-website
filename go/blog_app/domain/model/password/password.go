package password

import (
	"errors"
	"fmt"

	"golang.org/x/crypto/bcrypt"
)

var (
	ErrPasswordTooShort = errors.New("password must be at least 8 characters")
	ErrIncorrectPassword = errors.New("wrong password")
)

type Password struct {
	password       string
	hashedPassword string
}

func (p *Password) HashedPassword() string { return p.hashedPassword }

func validatePassword(p string) error {
	if len(p) < 8 {
		return ErrPasswordTooShort
	}
	return nil
}

func NewPassword(p string) (*Password, error) {
	if err := validatePassword(p); err != nil {
		return nil, err
	}
	hashedPassword, err := hash(p)
	if err != nil {
		return nil, err
	}

	return &Password{
		password:       p,
		hashedPassword: hashedPassword,
	}, nil
}

func Equals(hashedPassword, incomingPassword string) error {
	// return no error if hashed password and hashedPassword match
	if err := bcrypt.CompareHashAndPassword(
		[]byte(hashedPassword),
		[]byte(incomingPassword),
	); err != nil {
		return ErrIncorrectPassword
	}
	return nil
}

func hash(p string) (string, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(p), bcrypt.DefaultCost)
	if err != nil {
		return "", fmt.Errorf("failed to hash password: %v", err)
	}
	return string(hashedPassword), nil
}
