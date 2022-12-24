package repository

import
import "blog_app/domain/model"

type userRepository struct {
	db *db.Conn
}

func NewUserRepository(db *db.Conn) userRepository {
	return userRepository{}
}

func (r userRepository) CreateUser(user model.User) (model.User, error) {

}
