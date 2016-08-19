package admin

import (
	"github.com/zituocn/VMovie/models"
	"strconv"
	"strings"
)

type PageHandle struct {
	baseController
}

func (this *PageHandle) Add() {
	var info models.PageInfo
	this.Data["info"] = info
	this.Data["AdminDir"] = this.admindir
	this.TplName = "admin/pageadd.html"
}

func (this *PageHandle) Edit() {
	var (
		id   int64
		info models.PageInfo
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
	this.TplName = "admin/pageadd.html"
}

func (this *PageHandle) List() {
	var (
		page     int64
		pagesize int64 = 12
		offset   int64
		keyword  string
		pager    string
		list     []*models.PageInfo
		info     models.PageInfo
	)
	keyword = strings.TrimSpace(this.GetString("keyword"))
	pagestr := this.Ctx.Input.Param(":page")
	page, _ = strconv.ParseInt(pagestr, 10, 64)
	if page < 1 {
		page = 1
	}
	offset = (page - 1) * pagesize
	query := info.Query()
	if len(keyword) > 0 {
		query = query.Filter("name__icontains", keyword)
	}
	count, _ := query.Count()
	if count > 0 {
		query.OrderBy("-Id").Limit(pagesize, offset).All(&list)
	}
	pager = this.PageList(pagesize, page, count, false, this.admindir+"page/list")
	this.Data["pager"] = pager
	this.Data["list"] = list
	this.Data["admindir"] = this.admindir
	this.Data["keyword"] = keyword
	this.Data["count"] = count
	this.Data["AdminDir"] = this.admindir
	this.TplName = "admin/pagelist.html"
}

//保存专题
func (this *PageHandle) Save() {
	var (
		id          int64
		name        string
		ename       string
		content     string
		status      int64
		title       string
		keywords    string
		description string

		info models.PageInfo
		err  error
	)

	name = strings.TrimSpace(this.GetString("name"))
	ename = strings.TrimSpace(this.GetString("ename"))
	content = strings.TrimSpace(this.GetString("content"))
	title = strings.TrimSpace(this.GetString("title"))
	keywords = strings.TrimSpace(this.GetString("keywords"))
	description = strings.TrimSpace(this.GetString("description"))
	status, _ = this.GetInt64("status")
	id, _ = this.GetInt64("id")
	if len(name) == 0 || len(ename) == 0 || len(content) == 0 || len(title) == 0 || len(description) == 0 {
		this.showmsg("对不起,带*号的项必须填写...")
	}

	info.Name = name
	info.Ename = ename
	info.Editor = this.nickname
	info.Title = title
	info.Content = content
	info.Keywords = keywords
	info.Description = description
	info.Status = status
	if id > 0 {
		info.Id = id
		err = info.Update("name", "content", "title", "keywords", "description", "status")
	} else {
		err = info.Insert()
	}
	if err != nil {
		this.showmsg("保存出错，错误信息：" + err.Error())
	} else {
		this.showmsg("专题数据保存成功...", this.admindir+"page/add")
	}
}
