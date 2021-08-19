package controllers

import (
	"beego-hananoq-blog/component/errcode"
	"beego-hananoq-blog/conf"
)

type FileController struct {
	LoginController
}

func (c *FileController) OSSInfo() {
	defer func() {
		c.Output(errcode.SUCCESS)
	}()
	oss := conf.GetConfig().OSS
	data := map[string]interface{}{
		"accessKeyId":     oss.AccessKeyId,
		"accessKeySecret": oss.AccessKeySecret,
		"bucket":          oss.Bucket,
		"endpoint":        oss.Endpoint,
	}
	c.Resp["data"] = data
}
