package admin

import (
	//"github.com/astaxie/beego"
	"github.com/zituocn/VMovie/models"
	"strconv"
	"strings"
)

type MovieHandle struct {
	baseController
}

//add page
func (this *MovieHandle) Add() {
	//所有影片分类
	var (
		classlist  []*models.MovieClassInfo
		movieclass models.MovieClassInfo
		info       models.MovieInfo
	)
	movieclass.Query().All(&classlist)
	this.Data["classlist"] = classlist
	this.Data["AdminDir"] = this.admindir
	this.Data["info"] = info
	this.TplName = "admin/movieadd.html"
}

//edit page
func (this *MovieHandle) Edit() {
	var (
		id   int64
		info models.MovieInfo
	)

	//所有影片分类
	var classlist []*models.MovieClassInfo
	var movieclass models.MovieClassInfo
	movieclass.Query().All(&classlist)

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
	this.Data["classlist"] = classlist
	this.Data["AdminDir"] = this.admindir
	this.TplName = "admin/movieadd.html"
}

//list page
func (this *MovieHandle) List() {

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
		week     int64
		done     int64
		cid      int64
		status   int64 = -10
		pager    string
		err      error
	)
	keyword = strings.TrimSpace(this.GetString("keyword"))
	cid, _ = this.GetInt64("cid")
	week, _ = this.GetInt64("week")
	done, _ = this.GetInt64("done")
	status, err = this.GetInt64("status")
	if err != nil {
		status = -10
	}
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
	if week > 0 {
		query = query.Filter("updateweek", week)
	}
	if done == 1 {
		query = query.Filter("isend", 0)
	}
	if done == 2 {
		query = query.Filter("isend", 1)
	}
	if status != -10 {
		if status != 10 {
			query = query.Filter("status", status)
		} else {
			query = query.Filter("isnew", 1)
		}
	}
	count, _ := query.Count()
	if count > 0 {
		query.OrderBy("-Id").Limit(pagesize, offset).All(&list)
	}
	pager = this.PageList(pagesize, page, count, false, this.admindir+"movie/list")
	this.Data["pager"] = pager
	this.Data["list"] = list
	this.Data["admindir"] = this.admindir
	this.Data["keyword"] = keyword
	this.Data["cid"] = cid
	this.Data["week"] = week
	this.Data["done"] = done
	this.Data["status"] = status
	this.Data["count"] = count
	this.TplName = "admin/movielist.html"
}

//save post
func (this *MovieHandle) Save() {
	var (
		id       int64
		name     string
		ename    string
		actor    string
		director string
		writer   string
		language string
		content  string
		//tags        string
		title       string
		keywords    string
		description string
		playdate    string
		photo       string
		iphoto      string

		cid        int64 = 0
		status     int64 = 0
		episode    int64 = 0
		updateweek int64 = 0
		isnew      int64 = 0
		info       models.MovieInfo
		err        error
	)

	name = strings.TrimSpace(this.GetString("name"))
	ename = strings.TrimSpace(this.GetString("ename"))
	actor = strings.TrimSpace(this.GetString("actor"))
	director = strings.TrimSpace(this.GetString("director"))
	writer = strings.TrimSpace(this.GetString("writer"))
	language = strings.TrimSpace(this.GetString("language"))
	content = strings.TrimSpace(this.GetString("content"))

	//tags = strings.TrimSpace(this.GetString("tags"))
	title = strings.TrimSpace(this.GetString("title"))
	keywords = strings.TrimSpace(this.GetString("keywords"))
	description = strings.TrimSpace(this.GetString("description"))
	playdate = strings.TrimSpace(this.GetString("playdate"))
	photo = strings.TrimSpace(this.GetString("photo"))
	iphoto = strings.TrimSpace(this.GetString("iphoto"))

	id, _ = this.GetInt64("id")
	cid, _ = this.GetInt64("cid")
	status, _ = this.GetInt64("status")
	episode, _ = this.GetInt64("episode")
	updateweek, _ = this.GetInt64("updateweek")
	isnew, _ = this.GetInt64("isnew")

	if len(name) == 0 || len(photo) == 0 || len(ename) == 0 || len(actor) == 0 || len(director) == 0 || len(content) == 0 || cid == 0 || episode == 0 {
		this.showmsg("带*号的为必须填写的内容...")
	}

	info.Name = name
	info.Ename = ename
	info.Actor = actor
	info.Director = director
	info.Writer = writer
	info.Language = language
	info.Content = content
	info.Title = title
	info.Keywords = keywords
	info.Description = description

	info.Playdate = playdate
	info.Photo = photo
	info.Iphoto = iphoto

	info.Cid = cid
	info.Status = status
	info.Episode = episode
	info.Updateweek = updateweek
	info.Addtime = this.getTime()
	info.Updatetime = this.getTime()
	info.Editor = this.nickname
	info.Isnew = isnew //新剧推荐

	if id > 0 {
		info.Id = id
		err = info.Update("name", "ename", "actor", "director", "writer", "language", "content", "title", "keywords", "description", "playdate", "cid", "status", "Episode", "Updateweek", "photo", "iphoto", "Updatetime", "Isnew")
	} else {
		err = info.Insert()
	}
	if err != nil {
		this.showmsg("保存出错，错误信息：" + err.Error())
	} else {
		this.showmsg("数据保存成功...", this.admindir+"movie/add")
	}
}
