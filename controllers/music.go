package controllers

import "beego-hananoq-blog/component/errcode"

type MusicController struct {
	LoginController
}

func (c *MusicController) GetMusicList() {
	defer func() {
		c.Output(errcode.SUCCESS)
	}()
	c.Resp["data"] = map[string]interface{}{
		"list": []string{},
	}
}
