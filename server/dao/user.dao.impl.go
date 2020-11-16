package dao

import (
	"errors"
	"time"

	"lawencon.com/credential/model"
)

// UserDaoImpl for implement userdao
type UserDaoImpl struct{}

//Login for implement logindao
func (UserDaoImpl) Login(user model.UsersDb) (e error) {
	defer CatchError(&e)
	tx := g.First(&user)
	if tx.Error == nil {
		*user.UpdatedDate = model.Timestamp(time.Now())
		return g.Save(&user).Error
	}

	return errors.New("Invalid Login")
}

//ValidateToken for implement validatetokendao
func (UserDaoImpl) ValidateToken(token string) (e error) {
	defer CatchError(&e)
	var user model.UsersDb
	return g.Model(&model.UsersDb{}).Where("token", token).First(&user).Error
}

//Register for implement registerdao
func (UserDaoImpl) Register(user *model.UsersDb) (e error) {
	defer CatchError(&e)
	return g.Create(user).Error
}

// GetByUsername for implement getbyusernamedao
func (UserDaoImpl) GetByUsername(username string) (u model.UsersDb, e error) {
	defer CatchError(&e)
	var user = model.UsersDb{Username: username}
	tx := g.Where("username", username).Find(&user)
	return user, tx.Error
}

// UpdateUser for implement updateuserdao
func (UserDaoImpl) UpdateUser(user *model.UsersDb) (e error) {
	defer CatchError(&e)
	return g.Save(user).Error
}
