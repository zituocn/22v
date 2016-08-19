package admin

import (
	//"github.com/astaxie/beego"
	"github.com/zituocn/VMovie/models"
	"strconv"
	"strings"
)

type RecommendHandle struct {
	baseController
}

func (this *RecommendHandle) Add() {
	var info models.RecommendInfo
	this.Data["info"] = info
	this.TplName = "admin/recommendadd.html"
	this.Data["AdminDir"] = this.admindir
}

///保存
func (this *RecommendHandle) Save() {
	var (
		id      int64
		name    string
		content string
		info    models.RecommendInfo
		err     error
	)

	name = strings.TrimSpace(this.GetString("name"))
	content = strings.TrimSpace(this.GetString("content"))
	id, _ = this.GetInt64("id")
	if len(name) == 0 || len(content) == 0 {
		this.showmsg("对不起,带*号的项必须填写...")
	}

	info.Name = name
	info.Content = content
	if id > 0 {
		info.Id = id
		err = info.Update("name", "content")
	} else {
		err = info.Insert()
	}
	if err != nil {
		this.showmsg("保存出错，错误信息：" + err.Error())
	} else {
		this.showmsg("推荐数据保存成功...", this.admindir+"recommend/add")
	}
}

func (this *RecommendHandle) List() {
	var (
		page     int64
		pagesize int64 = 12
		offset   int64
		pager    string
		list     []*models.RecommendInfo
		info     models.RecommendInfo
	)
	pagestr := this.Ctx.Input.Param(":page")
	page, _ = strconv.ParseInt(pagestr, 10, 64)
	if page < 1 {
		page = 1
	}
	offset = (page - 1) * pagesize
	query := info.Query()
	count, _ := query.Count()
	if count > 0 {
		query.OrderBy("-Id").Limit(pagesize, offset).All(&list)
	}
	pager = this.PageList(pagesize, page, count, false, this.admindir+"recommend/list")
	this.Data["pager"] = pager
	this.Data["list"] = list
	this.Data["admindir"] = this.admindir
	this.Data["count"] = count

	this.TplName = "admin/recommendlist.html"
	this.Data["AdminDir"] = this.admindir
}

//editor page
func (this *RecommendHandle) Edit() {
	var (
		id   int64
		info models.RecommendInfo
	)
	idStr := this.Ctx.Input.Param(":id")
	id, _ = strconv.ParseInt(idStr, 10, 64)
	if id <= 0 {
		this.showmsg("数据错误，返回重试...")
	}
	info.Id = id
	err := info.Read()
	if err != nil {
		this.showmsg("数据不存在...")
	}

	this.Data["info"] = info
	this.Data["AdminDir"] = this.admindir
	this.TplName = "admin/recommendadd.html"
}
