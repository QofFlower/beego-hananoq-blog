package NSRouters

import (
	"beego-hananoq-blog/controllers"
	"github.com/astaxie/beego"
)

func init() {
	router := beego.NewNamespace("blog",
		beego.NSNamespace("auth",
			beego.NSRouter("/index", &controllers.MainController{}, "GET:Index"),
			beego.NSRouter("/token", &controllers.Oauth2Controller{}, "POST:Token"),
			beego.NSRouter("/login", &controllers.Oauth2Controller{}, "POST:Token"),
			beego.NSRouter("/noLogin", &controllers.Oauth2Controller{}, "GET:Logout"),
			beego.NSRouter("/check-token", &controllers.Oauth2Controller{}, "POST:CheckToken"),
			beego.NSRouter("/user-info", &controllers.Oauth2Controller{}, "GET:UserInfo"),
		),
		beego.NSNamespace("article",
			beego.NSRouter("/list", &controllers.ArticleController{}, "GET:List"),
			beego.NSRouter("/searchTag", &controllers.ArticleController{}, "GET:SearchTag"),
			beego.NSRouter("/save", &controllers.ArticleController{}, "POST:Save"),
			beego.NSRouter("/timeLine", &controllers.ArticleController{}, "GET:TimeLine"),
			beego.NSRouter("/:id", &controllers.ArticleController{}, "GET:SearchById"),
			beego.NSRouter("/search", &controllers.ArticleController{}, "GET:Search"),
		),
		beego.NSNamespace("friend-address",
			beego.NSRouter("/getFriendAddress", &controllers.FriendAddressController{}, "GET:GetFriendAddress"),
		),
		beego.NSNamespace("music",
			beego.NSRouter("/list", &controllers.MusicController{}, "GET:GetMusicList"),
		),
		beego.NSNamespace("file",
			beego.NSRouter("/oss-info", &controllers.FileController{}, "GET:OSSInfo"),
		),
	)
	beego.AddNamespace(router)
}
