package controllers

import (
	"beego-hananoq-blog/component/errcode"
	"beego-hananoq-blog/component/utils"
	"beego-hananoq-blog/oauth2"
	"beego-hananoq-blog/services/usersrv"
	"fmt"
	"github.com/astaxie/beego/logs"
	"strconv"
	"time"
)

type Oauth2Controller struct {
	LoginController
}

// Token 获取token
func (c *Oauth2Controller) Token() {
	errCode := errcode.SUCCESS
	defer func() {
		c.Output(errCode)
	}()
	if c.AccessToken != "" {
		if tokenInfo, err := oauth2.GetAccessToken(c.AccessToken); err == nil && tokenInfo != nil {
			c.Resp["data"] = tokenInfo
			return
		} else {
			logs.Error("err", err, tokenInfo)
		}
	}
	if tokenInfo, err := oauth2.GenerateToken(c.Ctx.Request); err != nil {
		errCode = errcode.INVALID_USERNAME_PASSWORD
		return
	} else {
		c.Resp["data"] = tokenInfo
	}
}

// CheckToken 检查token是否合法，并返回信息
func (c *Oauth2Controller) CheckToken() {
	defer func() {
		c.Output(errcode.SUCCESS)
	}()
	srv, manager := oauth2.GetAuhServer(), oauth2.GetManager()
	r := c.Ctx.Request
	token, err := srv.ValidationBearerToken(r)
	if err != nil {
		c.CustomOutput(errcode.ERROR, err.Error())
		c.StopRun()
		return
	}
	cli, err := manager.GetClient(token.GetClientID())
	if err != nil {
		c.CustomOutput(errcode.ERROR, err.Error())
		c.StopRun()
		return
	}
	fmt.Println(token)
	data := map[string]interface{}{
		"expires_in": int64(token.GetAccessCreateAt().Add(token.GetAccessExpiresIn()).Sub(time.Now()).Seconds()),
		"user_id":    token.GetUserID(),
		"client_id":  token.GetClientID(),
		"scope":      token.GetScope(),
		"domain":     cli.GetDomain(),
	}
	c.Resp["data"] = data
}

// Register 用户注册
func (c *Oauth2Controller) Register() {
	errCode := errcode.SUCCESS
	defer func() {
		c.Output(errCode)
	}()
	name, username, password := c.GetString("name"), c.GetString("username"), c.GetString("password")
	if username == "" || password == "" {
		errCode = errcode.USERNAME_PASSWORD_EMPTY
		return
	}
	if len(username) > 20 || len(password) > 20 {
		errCode = errcode.USERMELEN_PASSWORD_LENOUT
		return
	}
	if user, _ := usersrv.GetUserInfoByUsername(username); user != nil {
		errCode = errcode.USER_INFO_EXIST
		return
	}
	username, password = utils.Md5Encode(username), utils.Md5Encode(password)
	if errCode := usersrv.UserRegister(name, username, password); errCode != errcode.SUCCESS {
		return
	}
}

// Logout 用户登出
func (c *Oauth2Controller) Logout() {
	errCode := errcode.SUCCESS
	defer func() {
		c.Output(errCode)
	}()
	err := oauth2.GetManager().RemoveAccessToken(c.AccessToken)
	if err != nil {
		errCode = errcode.ERROR
		return
	}
}

// UserInfo 用户信息
func (c *Oauth2Controller) UserInfo() {
	errCode := errcode.SUCCESS
	defer func() {
		c.Output(errCode)
	}()
	if c.UserId == "" {
		errCode = errcode.INVALID_PARAM
		return
	}
	userId, err := strconv.Atoi(c.UserId)
	if err != nil {
		errCode = errcode.INVALID_PARAM
		return
	}
	user, code := usersrv.GetUserInfoById(userId)
	if code != errcode.SUCCESS {
		errCode = code
		return
	}
	c.Resp["data"] = user
}
