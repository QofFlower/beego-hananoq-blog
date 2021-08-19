package controllers

import "beego-hananoq-blog/component/errcode"

type MainController struct {
	LoginController
}

func (c *MainController) Get() {

}

func (c *MainController) Login() {
	defer func() {
		c.Output(errcode.SUCCESS)
	}()
}

func (c *MainController) Index() {
	c.CustomOutput(200, "Welcome to home page")
}
