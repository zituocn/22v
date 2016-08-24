package controllers

type HttpErrorHandle struct {
	baseController
}

func (this *HttpErrorHandle) Error404() {
	this.Data["content"] = "The Page is Not Found <br /> 对不起，页面不存在..."
	this.Data["title"] = "页面不存在"
	this.TplName = "_error.html"
}

func (this *HttpErrorHandle) Error501() {
	this.Data["content"] = "Server Error <br /> 对不起，服务器内部错误..."
	this.Data["title"] = "服务器错误"
	this.TplName = "_error.html"
}

func (this *HttpErrorHandle) ErrorDb() {
	this.Data["content"] = "DataBase Error <br /> 对不起，数据库错误..."
	this.Data["title"] = "数据库错误"
	this.TplName = "_error.html"
}
