package model

import (
	"music-saas/global"
)

type User struct {
	global.MODEL
	Username string `json:"username" gorm:"not null;type:varchar(50);unique_index;comment:用户登录名"`
	Password string `json:"-"  gorm:"not null;comment:用户登录密码"`
	NickName string `json:"nickName" gorm:"comment:用户昵称" `
	Avatar   string `json:"avatar" gorm:"default:https://cool-music-1257890402.cos.ap-nanjing.myqcloud.com/avatar/avatar.jpeg;comment:用户头像"`
	Age      int8   `json:"age" gorm:"comment:用户年龄"`
	Phone    string `json:"phone" gorm:"type:varchar(50);comment:手机号"`
	Email    string `json:"email" gorm:"type:varchar(50);comment:邮箱"`
	Sex      bool   `json:"sex" gorm:"comment:性别"`
	Status   bool   `json:"status" gorm:"not null;default:true;comment:状态 ACTIVE/DELETED"`
}

func (User) TableName() string {
	return "user"
}
