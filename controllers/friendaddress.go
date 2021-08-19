package controllers

import "beego-hananoq-blog/component/errcode"

type FriendAddressController struct {
	LoginController
}

type FriendAddress struct {
	Id      int    `json:"id"`
	Name    string `json:"name"`
	Address string `json:"address"`
}

func (c *FriendAddressController) GetFriendAddress() {
	defer func() {
		c.Output(errcode.SUCCESS)
	}()
	c.Resp["data"] = []*FriendAddress{
		{233, "哔哩哔哩", "https://www.bilibili.com"},
	}
}
