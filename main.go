package main

import (
	_ "beego-hananoq-blog/component/errcode"
	"beego-hananoq-blog/component/logger"
	"beego-hananoq-blog/component/requestlimit"
	_ "beego-hananoq-blog/conf"
	_ "beego-hananoq-blog/models"
	_ "beego-hananoq-blog/oauth2"
	_ "beego-hananoq-blog/routers"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context"
	"github.com/astaxie/beego/plugins/cors"
)

func init() {
	logger.InitLogger()
}

func main() {
	// 处理跨域请求，为什么这么配？参考https://www.cnblogs.com/dannyyao/p/8047319.html
	//beego.InsertFilter("*", beego.BeforeRouter, cors.Allow(&cors.Options{
	//	//AllowAllOrigins: true,
	//	AllowMethods:     []string{"*"},
	//	AllowHeaders:     []string{"Origin", "Authorization", "Access-Control-Allow-Origin", "Content-Type"},
	//	ExposeHeaders:    []string{"Content-Length", "Access-Control-Allow-Origin"},
	//	AllowCredentials: true,
	//	AllowOrigins:     []string{"https://*.*.*.*:*", "http://*.*.*.*:*"},
	//}))
	beego.InsertFilter("*", beego.BeforeRouter, cors.Allow(&cors.Options{
		AllowAllOrigins:  true,
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Authorization", "Access-Control-Allow-Origin", "Access-Control-Allow-Headers", "Content-Type"},
		ExposeHeaders:    []string{"Content-Length", "Access-Control-Allow-Origin", "Access-Control-Allow-Headers", "Content-Type"},
		AllowCredentials: true,
	}))
	beego.InsertFilter("*", beego.BeforeRouter, func(context *context.Context) {
		context.ResponseWriter.Header().Set("Access-Control-Allow-Origin", context.Request.Header.Get("Origin"))
	})
	// 限制请求此时的filter
	requestlimit.RunRate()
	beego.Run()
}
