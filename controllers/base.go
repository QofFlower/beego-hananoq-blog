package controllers

import (
	"beego-hananoq-blog/component/errcode"
	"beego-hananoq-blog/component/utils"
	"beego-hananoq-blog/oauth2"
	"github.com/astaxie/beego"
)

type BaseController struct {
	beego.Controller
	Resp map[string]interface{}
}

func (c *BaseController) Output(code uint) {
	c.Resp["code"] = code
	c.Resp["msg"] = errcode.GetMsg(code)

	c.Ctx.Input.SetData("respcode", c.Resp["code"])
	c.Ctx.Input.SetData("respmsg", c.Resp["msg"])

	c.Data["json"] = c.Resp
	c.ServeJSON()
}

func (c *BaseController) Options() {
	defer func() {
		c.Output(errcode.SUCCESS)
	}()
}

func (c *BaseController) CustomOutput(code uint, msg string) {
	c.Resp["code"] = code
	c.Resp["msg"] = msg
	c.Ctx.Input.SetData("respcode", c.Resp["code"])
	c.Ctx.Input.SetData("respmsg", c.Resp["msg"])

	c.Data["json"] = c.Resp
	c.ServeJSON()
}

type LoginController struct {
	BaseController
	AccessToken      string
	UserId           string
	UserNameNoEncode string
}

func (c *LoginController) Prepare() {
	c.Resp = make(map[string]interface{})
	c.AccessToken = c.GetString("access_token")
	c.UserId = c.GetString("user_id")
	timestamp := c.GetString("t")
	if timestamp == "" || !utils.IsNum(timestamp) {
		c.Output(errcode.INVALID_PARAM)
		c.StopRun()
	}
	requestUri := c.Ctx.Request.RequestURI
	if !utils.UriFilter(requestUri) && (c.AccessToken == "" || !c.CheckToken()) {
		c.Output(errcode.INVALID_TOKEN)
		c.StopRun()
	}
}

// CheckToken 验证token，请求任何接口时候调用
func (c *LoginController) CheckToken() bool {
	srv := oauth2.GetAuhServer()
	r := c.Ctx.Request
	if token, err := srv.ValidationBearerToken(r); token != nil && err == nil {
		c.UserId = token.GetUserID()
		c.AccessToken = token.GetAccess()
		return true
	} else {
		return false
	}
}
