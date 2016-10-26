package routers

import (
	//"fmt"
	"github.com/astaxie/beego"
	"github.com/zituocn/VMovie/controllers"
	"github.com/zituocn/VMovie/controllers/admin"
)

func init() {

	//pages
	beego.Router("/", &controllers.IndexHandle{}, "*:Index")
	beego.Router("/m/:cid:int/", &controllers.IndexHandle{}, "*:List")
	beego.Router("/m/:cid:int/:page:int/", &controllers.IndexHandle{}, "*:List")
	beego.Router("/v/:id:int/", &controllers.IndexHandle{}, "*:Detail")
	beego.Router("/search/:key(.+)/", &controllers.IndexHandle{}, "*:Search")
	beego.Router("/search/:key(.+)/:page:int/", &controllers.IndexHandle{}, "*:Search")
	beego.Router("/json/", &controllers.IndexHandle{}, "*:Json")
	beego.Router("/article/:id:int/", &controllers.IndexHandle{}, "*:Page")
	beego.Router("/today/", &controllers.IndexHandle{}, "*:Today")
	beego.Router("/new/", &controllers.IndexHandle{}, "*:New")
	beego.Router("/news/", &controllers.IndexHandle{}, "*:News")
	beego.Router("/news/:page:int/", &controllers.IndexHandle{}, "*:News")
	beego.Router("/22v.net.html", &controllers.IndexHandle{}, "*:Start")
	///error handel
	beego.ErrorController(&controllers.HttpErrorHandle{})

	//api json

	jns := beego.NewNamespace("/api",
		beego.NSRouter("/", &controllers.ApiHandle{}, "*:Index"),
		beego.NSRouter("/v/:id:int/", &controllers.ApiHandle{}, "*:Detail"),
		beego.NSRouter("/m/:cid:int/", &controllers.ApiHandle{}, "*:List"),
		beego.NSRouter("/m/:cid:int/:page:int/", &controllers.ApiHandle{}, "*:List"),
		beego.NSRouter("/search/:key(.+)/", &controllers.ApiHandle{}, "*:Search"),
		beego.NSRouter("/search/:key(.+)/:page:int/", &controllers.ApiHandle{}, "*:Search"),
		beego.NSRouter("/today/", &controllers.ApiHandle{}, "*:Today"),
		beego.NSRouter("/news/", &controllers.ApiHandle{}, "*:News"),            //资讯列表
		beego.NSRouter("/news/:page:int/", &controllers.ApiHandle{}, "*:News"),  //资讯翻页
		beego.NSRouter("/article/:id:int/", &controllers.ApiHandle{}, "*:Page"), //文章详情
		beego.NSRouter("/new/", &controllers.ApiHandle{}, "*:New"),              //新片列表
		beego.NSRouter("/new/:page:int/", &controllers.ApiHandle{}, "*:New"),    //新片列表
	)
	beego.AddNamespace(jns)

	///admin
	admindir := beego.AppConfig.String("admindir")
	if len(admindir) == 0 {
		admindir = "admin"
	}
	ns := beego.NewNamespace(admindir,
		beego.NSRouter("/", &admin.LoginHandle{}, "*:Login"),
		beego.NSRouter("/logout", &admin.LoginHandle{}, "*:Logout"),
		beego.NSRouter("/main", &admin.IndexHandle{}, "*:Main"),
		beego.NSRouter("/left", &admin.IndexHandle{}, "*:Left"),
		beego.NSRouter("/right", &admin.IndexHandle{}, "*:Right"),

		//影片管理
		beego.NSRouter("movie/add", &admin.MovieHandle{}, "*:Add"),
		beego.NSRouter("movie/edit/:id:int/", &admin.MovieHandle{}, "*:Edit"),
		beego.NSRouter("movie/save", &admin.MovieHandle{}, "post:Save"),
		beego.NSRouter("movie/list", &admin.MovieHandle{}, "*:List"),
		beego.NSRouter("movie/list/:page:int/", &admin.MovieHandle{}, "*:List"),

		//下载管理
		beego.NSRouter("down/add/:mid:int/", &admin.DownaddrHandle{}, "*:Add"),
		beego.NSRouter("down/list", &admin.DownaddrHandle{}, "*:List"),
		beego.NSRouter("down/save/:ep:int/", &admin.DownaddrHandle{}, "*:Save"),

		//图片上传管理
		beego.NSRouter("upload/add", &admin.UploadHandle{}, "*:UpLoadPage"),
		beego.NSRouter("upload/qiniusave", &admin.UploadHandle{}, "*:QiniuUpLoadFile"),

		//用户
		beego.NSRouter("user/changepassword", &admin.UserHandle{}, "*:ChangePass"),
		beego.NSRouter("user/savepass", &admin.UserHandle{}, "*:SavePass"),

		//影片关系
		beego.NSRouter("relation/add", &admin.RelationHandle{}, "*:Add"),
		beego.NSRouter("relation/save", &admin.RelationHandle{}, "*:Save"),
		beego.NSRouter("relation/list", &admin.RelationHandle{}, "*:List"),
		beego.NSRouter("relation/list/:page:int/", &admin.RelationHandle{}, "*:List"),
		beego.NSRouter("relation/detail/:id:int/", &admin.RelationHandle{}, "*:Detail"),
		beego.NSRouter("relation/delete/:id:int/", &admin.RelationHandle{}, "*:Delete"),

		//资讯管理
		beego.NSRouter("page/add", &admin.PageHandle{}, "*:Add"),
		beego.NSRouter("page/edit/:id:int/", &admin.PageHandle{}, "*:Edit"),
		beego.NSRouter("page/save", &admin.PageHandle{}, "*:Save"),
		beego.NSRouter("page/list", &admin.PageHandle{}, "*:List"),
		beego.NSRouter("page/list/:page:int/", &admin.PageHandle{}, "*:List"),

		//推荐管理
		beego.NSRouter("recommend/add", &admin.RecommendHandle{}, "*:Add"),
		beego.NSRouter("recommend/list", &admin.RecommendHandle{}, "*:List"),
		beego.NSRouter("recommend/list/:page:int/", &admin.RecommendHandle{}, "*:List"),
		beego.NSRouter("recommend/edit/:id:int/", &admin.RecommendHandle{}, "*:Edit"),
		beego.NSRouter("recommend/save", &admin.RecommendHandle{}, "*:Save"),
	)
	beego.AddNamespace(ns)
}
