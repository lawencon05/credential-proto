package service

import (
	"errors"
	"time"

	"golang.org/x/crypto/bcrypt"
	"lawencon.com/credential/dao"
	"lawencon.com/credential/model"
)

// UserServiceImpl for implement userservice
type UserServiceImpl struct{}

var userDao dao.UserDao = dao.UserDaoImpl{}

//Login for implement loginservice
func (UserServiceImpl) Login(user *model.UsersDb) (e error) {
	defer CatchError(&e)
	result, err := userDao.GetByUsername(user.Username)
	if err == nil {
		var err = bcrypt.CompareHashAndPassword([]byte(result.Password), []byte(user.Password))
		if err == nil {
			t, err := GenerateToken(user.Username)
			if err == nil {
				result.Token = t
				user.Token = t

				updateTime := model.Timestamp(time.Now())
				result.UpdatedDate = &updateTime
				err = userDao.UpdateUser(&result)
				return err
			}
		}
	}
	return errors.New("Invalid Username/Password")
}

//ValidateToken for implement validatetokenservice
func (UserServiceImpl) ValidateToken(token string) (e error) {
	defer CatchError(&e)
	return userDao.ValidateToken(token)
}

//Register for implement registerservice
func (UserServiceImpl) Register(user *model.UsersDb) (e error) {
	defer CatchError(&e)
	result, err := bcrypt.GenerateFromPassword([]byte(user.Password), 4)
	if err == nil {
		user.CreatedDate = model.Timestamp(time.Now())
		user.Password = string(result)
		return userDao.Register(user)
	}
	return err
}
