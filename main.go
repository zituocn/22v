package main

import (
	"github.com/astaxie/beego"
	"github.com/zituocn/VMovie/models"
	_ "github.com/zituocn/VMovie/routers"
)

func main() {
	beego.AddFuncMap("getclassname", models.GetMovieClassNameByCid)
	beego.AddFuncMap("GetMovieUpdateEP", models.GetMovieUpdateEP)
	beego.AddFuncMap("GetIPhoto", models.GetIPhoto)
	beego.AddFuncMap("GetMovieUpdateEPString", models.GetMovieUpdateEPString)
	beego.Run()
}
