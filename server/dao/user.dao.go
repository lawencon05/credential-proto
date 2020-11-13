package dao

import "lawencon.com/credential/model"

// UserDao for userdao contract
type UserDao interface {
	Login(user model.UsersDb) error
	ValidateToken(token string) error
	Register(user *model.UsersDb) error
	GetByUsername(username string) (model.UsersDb, error)
}
