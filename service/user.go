package service

import (
	"errors"
	"gorm.io/gorm"
	"music-saas/global"
	"music-saas/model"
	"music-saas/utils"
)

func Register(u model.User) (user model.User, err error) {
	var userModel model.User
	if !errors.Is(global.DB.Where("username = ?", u.Username).First(&userModel).Error, gorm.ErrRecordNotFound) {
		return user, errors.New("the username already exists")
	}
	u.Password = utils.Md5Encode([]byte(u.Password))
	err = global.DB.Create(&u).Error
	return u, err
}

func Login(u model.User) (user model.User, err error) {
	var userModel model.User
	u.Password = utils.Md5Encode([]byte(u.Password))
	err = global.DB.Where("username = ? AND password = ?", u.Username, u.Password).First(&userModel).Error
	return userModel, err
}

//@function: FindUserById
//@description: 通过id获取用户信息
//@param: id uint
//@return: err error, user *model.SysUser
func FindUserById(id uint) (user model.User, err error) {
	err = global.DB.Where("`id` = ?", id).First(&user).Error
	return user, err
}
