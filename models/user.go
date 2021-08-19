package models

import (
	"fmt"
	"github.com/astaxie/beego/logs"
	"github.com/astaxie/beego/orm"
	"time"
)

const (
	UserStatusAvailable = 1
	UserStatusBan       = 0

	UserTypeAdmin = iota
	UserTypeUser
)

type User struct {
	Id          int       `json:"id"`
	Name        string    `json:"name"`
	Username    string    `json:"username"`
	Password    string    `json:"password,omitempty"`
	AgiPassword string    `json:"agi_password,omitempty"`
	HeadPic     string    `json:"head_pic"`
	CreateTime  time.Time `json:"create_time"`
	UpdateTime  time.Time `json:"update_time"`
	Status      int       `json:"status"`
	Type        int       `json:"type"`
}

// GetUserIdByUsernameAndPwd 根据用户名和密码获取userId
func GetUserIdByUsernameAndPwd(username, password string) (userId string) {
	sql := `select id from hananoq_blog.blog_user where username = ? and password = ?`
	logs.Debug("Select by id and password SQL: ", sql)
	o := orm.NewOrm()
	if err := o.Raw(sql, username, password).QueryRow(&userId); err != nil {
		logs.Error("select error: ", err)
		return ""
	}
	return
}

// InsertUserInfo 插入用户信息
func InsertUserInfo(name, username, password string) (err error) {
	sql := `insert into hananoq_blog.h_user(%s) values()`
	cols := `name, username, password, create_time, update_time, status, type`
	values := []interface{}{name, username, password, time.Now(), time.Now(), UserStatusAvailable, UserTypeUser}
	o := orm.NewOrm()
	sql = fmt.Sprintf(sql, cols)
	logs.Debug("Insert user SQL: ", sql)
	if _, err = o.Raw(sql, values).Exec(); err != nil {
		return
	}
	return
}

// GetUserInfoByUsername 根据username获取用户信息
func GetUserInfoByUsername(username string) (user *User) {
	sql := `select * from hananoq_blog.blog_user where username = ?`
	o := orm.NewOrm()
	if err := o.Raw(sql, username).QueryRow(&user); err != nil {
		return
	}
	return
}

func GetUserInfoById(userId int) (user *User) {
	sql := `select %s from hananoq_blog.blog_user where id = ?`
	cols := `id, name, username, head_pic, create_time, update_time, status, type`
	o := orm.NewOrm()
	sql = fmt.Sprintf(sql, cols)
	if err := o.Raw(sql, userId).QueryRow(&user); err != nil {
		return
	}
	return
}
