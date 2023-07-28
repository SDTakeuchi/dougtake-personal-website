package password

import (
	"errors"
	"fmt"

	"golang.org/x/crypto/bcrypt"
)

const MinPasswordLength = 8

var (
	ErrPasswordTooShort  = errors.New("password must be at least 8 characters")
	ErrIncorrectPassword = errors.New("wrong password")
)

type Password struct {
	password       string
	hashedPassword string
}

// String() is an alias for Password.HashedPassword()
func (p Password) String() string { return p.HashedPassword() }

func (p Password) HashedPassword() string { return p.hashedPassword }

func validatePassword(p string) error {
	if len(p) < MinPasswordLength {
		return ErrPasswordTooShort
	}
	return nil
}

func NewPassword(rawPassword string) (*Password, error) {
	if err := validatePassword(rawPassword); err != nil {
		return nil, err
	}
	hashedPassword, err := hash(rawPassword)
	if err != nil {
		return nil, err
	}

	return &Password{
		password:       rawPassword,
		hashedPassword: hashedPassword,
	}, nil
}

func ParseHashedPassword(hp string) *Password {
	return &Password{
		"",
		hp,
	}
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
