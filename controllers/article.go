package controllers

import (
	"beego-hananoq-blog/component/errcode"
	"beego-hananoq-blog/models"
	"beego-hananoq-blog/services/articlesrv"
	"beego-hananoq-blog/services/usersrv"
	"encoding/json"
	"strconv"
)

type ArticleController struct {
	LoginController
}

// List 博客列表
func (c *ArticleController) List() {
	errCode := errcode.SUCCESS
	defer func() {
		c.Output(errCode)
	}()
	pageSize, err1 := c.GetInt("pageSize")
	pageIndex, err2 := c.GetInt("currentPage")
	if err1 != nil || err2 != nil || pageSize <= 0 || pageIndex <= 0 {
		errCode = errcode.INVALID_PARAM
	}
	data := make(map[string]interface{})
	data["total"], data["list"] = articlesrv.GetArticleList(pageSize, pageIndex)
	c.Resp["data"] = data
}

// SearchTag 标签内容墙
func (c *ArticleController) SearchTag() {
	defer func() {
		c.Output(errcode.SUCCESS)
	}()
	c.Resp["data"] = articlesrv.SearchTag()
}

// Save 发表文章
func (c *ArticleController) Save() {
	errCode := errcode.SUCCESS
	defer func() {
		c.Output(errCode)
	}()
	var article *models.Article
	if err := json.Unmarshal(c.Ctx.Input.RequestBody, &article); err != nil {
		errCode = errcode.INVALID_JSON_BODY
		return
	}
	user, _ := usersrv.GetUserInfoById(article.ManagerId)
	if user.Type != models.UserTypeAdmin {
		errCode = errcode.PERMISSION_DENY
		return
	}
	article.ManagerName = user.Name
	articlesrv.InsertArticle(article)
}

// TimeLine 获取时间线
func (c *ArticleController) TimeLine() {
	errCode := errcode.SUCCESS
	defer func() {
		c.Output(errCode)
	}()
	c.Resp["data"] = articlesrv.GetArticleListOrderByCreateTime()
}

// SearchById 根据id查找
func (c *ArticleController) SearchById() {
	errCode := errcode.SUCCESS
	defer func() {
		c.Output(errCode)
	}()
	str := c.Ctx.Input.Param(":id")
	id, err := strconv.Atoi(str)
	if err != nil {
		errCode = errcode.INVALID_PARAM
	}
	c.Resp["data"] = articlesrv.GetById(id)
}

func (c *ArticleController) Search() {
	errCode := errcode.SUCCESS
	defer func() {
		c.Output(errCode)
	}()
	keywords := c.GetString("keywords")
	pageIndex, err := c.GetInt("currentPage")
	if err != nil {
		errCode = errcode.INVALID_PARAM
		return
	}
	pageSize, err := c.GetInt("pageSize")
	if err != nil {
		errCode = errcode.INVALID_PARAM
		return
	}
	data := make(map[string]interface{})
	data["total"], data["list"] = articlesrv.Search(keywords, pageIndex, pageSize)
	c.Resp["data"] = data
}
