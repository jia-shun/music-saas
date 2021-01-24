package request

type Register struct {
	Username string `json:"username"`
	Password string `json:"password"`
	NickName string `json:"nickName" gorm:"comment:用户昵称" `
	Avatar   string `json:"avatar" gorm:"default:https://cool-music-1257890402.cos.ap-nanjing.myqcloud.com/avatar/avatar.jpeg"`
	Age      int8   `json:"age"`
	Phone    string `json:"phone"`
	Email    string `json:"email"`
	Sex      bool   `json:"sex"`
}

type Login struct {
	Username string `json:"username"`
	Password string `json:"password"`
}
