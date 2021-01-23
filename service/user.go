package service

import (
	"music/global"
	"music/model"
)

//@function: FindUserById
//@description: 通过id获取用户信息
//@param: id int
//@return: err error, user *model.SysUser
func FindUserById(id int) (err error, user *model.User) {
	var u model.User
	err = global.DB.Where("`id` = ?", id).First(&u).Error
	return err, &u
}
