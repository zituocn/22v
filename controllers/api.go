package controllers

import (
	//"encoding/json"
	//"fmt"
	//"github.com/astaxie/beego/orm"
	"github.com/zituocn/VMovie/models"
	"strconv"
	"strings"
	//"time"
)

type ApiHandle struct {
	baseController
}

//资讯文章列表
func (this *ApiHandle) News() {
	this.Ctx.Output.Header("Cache-Control", "public")
	var (
		page     int64
		pagesize int64 = 12
		offset   int64
		list     []*models.PageInfo
		info     models.PageInfo
		out      *models.ApiPageListInfo = new(models.ApiPageListInfo)
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
	out.List = list
	out.Page = page
	out.RecordCount = count

	this.Data["json"] = out
	this.ServeJSON()
}

//资讯文章详情
func (this *ApiHandle) Page() {
	var (
		out  *models.ApiPageDetailInfo = new(models.ApiPageDetailInfo)
		info *models.PageInfo          = new(models.PageInfo)
		id   int64
		err  error
	)
	//页面cache控制
	this.Ctx.Output.Header("Cache-Control", "public")

	idstr := this.Ctx.Input.Param(":id")
	id, err = strconv.ParseInt(idstr, 10, 64)
	if err != nil || id <= 0 {
		this.Abort("404")
		return
	}

	//读取数据
	info.Id = id
	err = info.Read()
	if err != nil || info.Status < 0 {
		this.Abort("404")
		return
	}

	//更新点击
	info.Views++
	info.Update("Views")

	out.Info = info
	this.Data["json"] = out
	this.ServeJSON()
}

//新剧推荐列表
func (this *ApiHandle) New() {
	var (
		info *models.MovieInfo
		list []*models.MovieInfo
		out  *models.ApiNewInfo = new(models.ApiNewInfo)
	)
	this.Ctx.Output.Header("Cache-Control", "public")

	info.Query().Filter("isnew", 1).OrderBy("-Updatetime").All(&list, "id", "name", "ename", "photo", "cid", "Language", "Updateweek", "Playdate", "isend", "Episode", "hasepisode")

	out.List = list
	for i := 0; i < len(list); i++ {
		list[i].Cname = models.GetMovieClassNameByCid(list[i].Cid)
	}
	this.Data["json"] = out
	this.ServeJSON()
}

//今日更新的影片
func (this *ApiHandle) Today() {
	var (
		info models.MovieInfo
		out  *models.ApiTodayInfo = new(models.ApiTodayInfo)
	)
	this.Ctx.Output.Header("Cache-Control", "public")
	list := info.GetWeekList(100)
	out.List = list
	for i := 0; i < len(list); i++ {
		list[i].Cname = models.GetMovieClassNameByCid(list[i].Cid)
	}
	this.Data["json"] = out
	this.ServeJSON()
}

//搜索api
func (this *ApiHandle) Search() {
	var (
		keyword  string
		page     int64
		pagesize int64 = 10
		offset   int64
		info     models.MovieInfo
		list     []*models.MovieInfo
		out      *models.ApiSearchInfo = new(models.ApiSearchInfo)
	)
	this.Ctx.Output.Header("Cache-Control", "public")
	keyword = strings.TrimSpace(this.Ctx.Input.Param(":key"))
	pagestr := this.Ctx.Input.Param(":page")
	page, _ = strconv.ParseInt(pagestr, 10, 64)
	if page < 1 {
		page = 1
	}
	offset = (page - 1) * pagesize

	query := info.Query()
	query = query.Filter("status__gte", 0)
	query = query.Filter("title__icontains", keyword).OrderBy("-id")

	count, _ := query.Count()
	if count > 0 {
		query.Limit(pagesize, offset).All(&list)
	}

	for i := 0; i < len(list); i++ {
		list[i].Cname = models.GetMovieClassNameByCid(list[i].Cid)
	}
	out.RecordCount = count
	out.Page = page
	out.List = list

	this.Data["json"] = out
	this.ServeJSON()
}

///首页api输出
func (this *ApiHandle) Index() {
	var (
		out    *models.ApiIndexInfo = new(models.ApiIndexInfo)
		info   models.MovieInfo
		mphoto []*models.MovieInfo
	)

	info.Query().Filter("status", 1).Limit(10, 0).OrderBy("-Updatetime").All(&mphoto, "id", "name", "ename", "photo", "cid", "Language", "Updateweek", "Playdate", "isend", "Episode", "hasepisode")

	var data1 models.ApiIndexMovieList
	data1.MList = info.GetWeekList(10)
	for i := 0; i < len(data1.MList); i++ {
		data1.MList[i].Cname = models.GetMovieClassNameByCid(data1.MList[i].Cid)
	}
	out.List = append(out.List, &data1)

	var data2 models.ApiIndexMovieList
	data2.MList = mphoto
	for i := 0; i < len(data2.MList); i++ {
		data2.MList[i].Cname = models.GetMovieClassNameByCid(data2.MList[i].Cid)
	}
	out.List = append(out.List, &data2)

	this.Data["json"] = out
	this.ServeJSON()
}

///列表页api输出
func (this *ApiHandle) List() {
	var (
		cid      int64
		page     int64
		pagesize int64 = 10
		offset   int64
		info     models.MovieInfo
		list     []*models.MovieInfo
		cinfo    *models.MovieClassInfo = new(models.MovieClassInfo)
		out      *models.ApiListInfo    = new(models.ApiListInfo) //输出模型
		err      error
	)

	cidstr := this.Ctx.Input.Param(":cid")
	cid, err = strconv.ParseInt(cidstr, 10, 64)
	if err != nil || cid <= 0 {
		this.Abort("404")
	}

	//查询分类信息
	cinfo.Id = cid
	err = cinfo.Read()

	pagestr := this.Ctx.Input.Param(":page")
	page, _ = strconv.ParseInt(pagestr, 10, 64)
	if page < 1 {
		page = 1
	}
	offset = (page - 1) * pagesize

	query := info.Query()
	query = query.Filter("status__gte", 0)
	if cid > 0 && cid != 100 && cid != 200 {
		query = query.Filter("cid", cid).OrderBy("-id")
	} else if cid == 100 {
		query = query.OrderBy("-id")
	} else if cid == 200 {
		query = query.OrderBy("-views")
	}

	count, _ := query.Count()
	if count > 0 {
		query.Limit(pagesize, offset).All(&list)
	}

	for i := 0; i < len(list); i++ {
		list[i].Cname = models.GetMovieClassNameByCid(list[i].Cid)
	}

	out.Cinfo = cinfo
	out.Page = page
	out.RecordCount = count
	out.MList = list
	this.Data["json"] = out
	this.ServeJSON()
}

///详情页api输出
func (this *ApiHandle) Detail() {
	var (
		out    *models.ApiDetailInfo = new(models.ApiDetailInfo)
		id     int64
		info   *models.MovieInfo   = new(models.MovieInfo)
		rmlist []*models.MovieInfo //相关影片数据
		rinfo  models.RelationInfo //影片关系
		down   models.DownAddrInfo
		cinfo  *models.MovieClassInfo = new(models.MovieClassInfo) //分类信息
		err    error
	)

	//get参
	idstr := this.Ctx.Input.Param(":id")
	id, err = strconv.ParseInt(idstr, 10, 64)

	if err != nil || id <= 0 {
		this.Abort("404")
		return
	}

	//读取数据
	info.Id = id
	err = info.Read()
	if err != nil || info.Status < 0 {
		this.Abort("404")
		return
	}

	//分类信息
	cinfo.Id = info.Cid
	cinfo.Read()

	//相关影片
	query := rinfo.Query().Filter("mids__icontains", ","+idstr+",")
	query.OrderBy("-Id").One(&rinfo)

	ids := make([]int64, 0)
	midstr := strings.Split(rinfo.Mids, ",")
	for _, s := range midstr {
		i, _ := strconv.ParseInt(s, 10, 64)
		if i > 0 && i != id {
			ids = append(ids, i)
		}
	}
	rmlist = make([]*models.MovieInfo, 0)
	if len(ids) > 0 {
		q := info.Query().Filter("id__in", ids)
		count, _ := q.Count()
		if count > 0 {
			q.OrderBy("-Id").Limit(10, 0).All(&rmlist, "Id", "Name", "Ename")
		}
	}

	//下载地址json数据
	list := make([]*models.DownAddrInfo, 0)
	down.Query().Filter("mid", id).OrderBy("ep").All(&list)

	//更新点击
	info.Views++
	info.Update("Views")

	out.Minfo = info
	out.SameList = rmlist
	out.DownList = list
	out.Cinfo = cinfo

	// json输出
	this.Data["json"] = out
	this.ServeJSON()

}
