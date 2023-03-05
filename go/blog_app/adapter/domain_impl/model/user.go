package model

import (
	"blog_app/adapter/persistance/database/postgres"
	"blog_app/domain/model"
	"blog_app/domain/model/password"
	"blog_app/domain/model/uuid"
)

type user struct {
	id       uuid.UUID
	name     string
	email    string
	password password.Password
}

func (t *user) ID() uuid.UUID { return t.id }

func (t *user) Name() string { return t.name }

func (t *user) Email() string { return t.email }

func (t *user) Password() password.Password { return t.password }

func NewUser(
	name string,
	email string,
	password password.Password,
) (model.User, error) {
	u := &user{
		name:     name,
		email:    email,
		password: password,
	}
	if err := model.ValidateUser(u); err != nil {
		return nil, err
	}
	return u, nil
}

func UserFromRecord(r postgres.User) model.User {
	return &user{
		id:       r.ID,
		name:     r.Name,
		email:    r.Email,
		password: *password.ParseHashedPassword(r.Password),
	}
}

func UserToRecord(m model.User) postgres.User {
	return postgres.User{
		ID:       m.ID(),
		Name:     m.Name(),
		Email:    m.Email(),
		Password: m.Password().String(),
	}
}
