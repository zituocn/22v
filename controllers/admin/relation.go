package admin

import (
	//"github.com/astaxie/beego"
	"github.com/zituocn/VMovie/models"
	"strconv"
	"strings"
	//"fmt"
)

type RelationHandel struct {
	baseController
}

type ErrorInfo struct {
	Code int64
	Msg  string
}

func (this *RelationHandel) Delete() {
	var (
		id    int64
		info  models.RelationInfo
		einfo ErrorInfo
		err   error
	)
	idstr := this.Ctx.Input.Param(":id")
	id, err = strconv.ParseInt(idstr, 10, 64)
	if err != nil || id <= 0 {
		einfo.Code = -1
		einfo.Msg = "参数错误，请重试..."
	}
	info.Id = id
	err = info.Delete()
	if err != nil {
		einfo.Code = -1
		einfo.Msg = "删除失败，请重试..."
	} else {
		einfo.Code = 0
		einfo.Msg = "删除成功..."
	}

	this.Data["json"] = &einfo
	this.ServeJSON()
}

func (this *RelationHandel) Detail() {
	var (
		id       int64
		page     int64
		pagesize int64 = 12
		offset   int64
		list     []*models.MovieInfo
		movie    models.MovieInfo
		info     models.RelationInfo
		pager    string
		err      error
	)

	idstr := this.Ctx.Input.Param(":id")
	id, err = strconv.ParseInt(idstr, 10, 64)
	if err != nil || id <= 0 {
		this.showmsg("参数错误，请返回重试...")
		return
	}

	pagestr := this.Ctx.Input.Param(":page")
	page, _ = strconv.ParseInt(pagestr, 10, 64)
	if page < 1 {
		page = 1
	}

	//取出关系信息
	info.Id = id
	err = info.Read()
	if err != nil {
		this.showmsg("此影片关系不存在...")
		return
	}
	offset = (page - 1) * pagesize
	//in 查询数据库
	query := movie.Query()
	ids := make([]int64, 0)
	midstr := strings.Split(info.Mids, ",")
	for _, s := range midstr {
		i, _ := strconv.ParseInt(s, 10, 64)
		if i > 0 {
			ids = append(ids, i)
		}
	}

	if len(info.Mids) > 0 {
		query = query.Filter("id__in", ids)
	}
	count, _ := query.Count()
	if count > 0 {
		query.OrderBy("-Id").Limit(pagesize, offset).All(&list)
	}
	pager = this.PageList(pagesize, page, count, false, this.admindir+"relation/add")
	this.Data["pager"] = pager
	this.Data["list"] = list
	this.Data["admindir"] = this.admindir
	this.Data["count"] = count
	this.Data["info"] = info
	this.TplName = "admin/relationdetail.html"
}

///影片关系列表
func (this *RelationHandel) List() {
	var (
		page     int64
		pagesize int64 = 12
		offset   int64
		keyword  string
		pager    string
		list     []*models.RelationInfo
		info     models.RelationInfo
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
	pager = this.PageList(pagesize, page, count, false, this.admindir+"relation/list")
	this.Data["pager"] = pager
	this.Data["list"] = list
	this.Data["admindir"] = this.admindir
	this.Data["keyword"] = keyword
	this.Data["count"] = count
	this.TplName = "admin/relationlist.html"
}

//后台保存关系
func (this *RelationHandel) Save() {
	var (
		mids  string
		name  string
		err   error
		rinfo models.RelationInfo
	)
	mids = strings.TrimSpace(this.GetString("mids"))
	name = strings.TrimSpace(this.GetString("name"))
	if len(mids) == 0 {
		this.showmsg("没有选择影片，保存失败...")
		return
	}
	if len(name) == 0 {
		this.showmsg("没有输入关系名称，保存失败...")
		return
	}

	rinfo.Mids = "," + mids + ","
	rinfo.Name = name
	query := rinfo.Query()
	query = query.Filter("mids", mids)
	count, _ := query.Count()
	if count > 0 {
		this.showmsg("此影片关系已经存在，保存失败...")
	} else {
		//写入新数据
		err = rinfo.Insert()
		if err != nil {
			this.showmsg("保存出错，错误信息：" + err.Error())
		} else {
			this.showmsg("影片关系保存成功....", this.admindir+"relation/add")
		}
	}

}

//add page
func (this *RelationHandel) Add() {
	//所有影片分类
	var classlist []*models.MovieClassInfo
	var movieclass models.MovieClassInfo
	movieclass.Query().All(&classlist)
	this.Data["classlist"] = classlist

	var (
		page     int64
		pagesize int64 = 12
		offset   int64
		list     []*models.MovieInfo
		movie    models.MovieInfo
		keyword  string
		cid      int64
		pager    string
	)
	keyword = this.GetString("keyword")
	cid, _ = this.GetInt64("cid")
	pagestr := this.Ctx.Input.Param(":page")
	page, _ = strconv.ParseInt(pagestr, 10, 64)
	if page < 1 {
		page = 1
	}
	offset = (page - 1) * pagesize
	query := movie.Query()
	if len(keyword) > 0 {
		query = query.Filter("title__icontains", keyword)
	}
	if cid > 0 {
		query = query.Filter("cid", cid)
	}

	count, _ := query.Count()
	if count > 0 {
		query.OrderBy("-Id").Limit(pagesize, offset).All(&list)
	}
	pager = this.PageList(pagesize, page, count, false, this.admindir+"relation/add")
	this.Data["pager"] = pager
	this.Data["list"] = list
	this.Data["admindir"] = this.admindir
	this.Data["keyword"] = keyword
	this.Data["cid"] = cid
	this.Data["count"] = count
	this.TplName = "admin/relationadd.html"
}
