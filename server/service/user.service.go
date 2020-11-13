package service

import "lawencon.com/credential/model"

// UserService for userservice contract
type UserService interface {
	Login(user *model.UsersDb) error
	ValidateToken(token string) error
	Register(user *model.UsersDb) error
}
