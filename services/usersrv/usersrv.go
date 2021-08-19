package usersrv

import (
	"beego-hananoq-blog/component/errcode"
	"beego-hananoq-blog/models"
	"github.com/astaxie/beego/logs"
)

// UserRegister 用户注册
func UserRegister(name, username, password string) (errCode uint) {
	userID := models.GetUserIdByUsernameAndPwd(username, password)
	if userID != "" {
		errCode = errcode.USER_INFO_EXIST
		return
	}
	if err := models.InsertUserInfo(name, username, password); err != nil {
		logs.Error("Insert User Error:", err.Error())
		return
	}
	return
}

func GetUserInfoByUsername(username string) (user *models.User, errCode uint) {
	return models.GetUserInfoByUsername(username), errcode.SUCCESS
}

func GetUserInfoById(userId int) (user *models.User, errCode uint) {
	return models.GetUserInfoById(userId), errcode.SUCCESS
}
