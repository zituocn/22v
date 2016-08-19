package admin

import (
	//"github.com/astaxie/beego"
	"github.com/zituocn/VMovie/models"
	"strconv"
	//"strings"
	"encoding/json"
	"fmt"
)

type DownaddrHandel struct {
	baseController
}

//add downaddr page
func (this *DownaddrHandel) Add() {
	var (
		mid      int64
		movie    models.MovieInfo
		down     models.DownAddrInfo
		downlist string
	)
	idStr := this.Ctx.Input.Param(":mid")
	mid, _ = strconv.ParseInt(idStr, 10, 64)
	if mid <= 0 {
		this.showmsg("数据错误，请返回重试...")
		return
	}
	movie.Id = mid
	err := movie.Read()
	if err != nil {
		this.showmsg("影片不存在，请返回重试...")
		return
	}
	list := make([]*models.DownAddrInfo, 0)
	down.Query().Filter("mid", mid).OrderBy("ep").All(&list)
	if b, err := json.Marshal(list); err == nil {
		downlist = string(b)
	}
	this.Data["downlist"] = downlist
	this.Data["episode"] = movie.Episode
	this.Data["movie"] = movie
	this.Data["AdminDir"] = this.admindir
	this.TplName = "admin/downadd.html"
}

//保存下载 post
func (this *DownaddrHandel) Save() {
	var (
		ep      int64
		mid     int64
		episode int64
		hdtv    string
		mkv     string
		info    models.DownAddrInfo
		minfo   models.MovieInfo
	)
	epStr := this.Ctx.Input.Param(":ep")
	ep, _ = strconv.ParseInt(epStr, 10, 64)
	hdtv = this.GetString("hdtv")
	mkv = this.GetString("mkv")
	mid, _ = this.GetInt64("mid")
	episode, _ = this.GetInt64("episode")

	if ep <= 0 {
		this.showmsg("数据错误，请返回重试...")
		return
	}

	if mid <= 0 || episode <= 0 {
		this.showmsg("数据错误，请返回重试...")
		return
	}

	info.Ep = ep
	info.Mid = mid
	err := info.Read("mid", "ep")
	///取出影片信息
	minfo.Id = mid
	minfo.Read()
	if err == nil {
		info.Hdtvurl = hdtv
		info.Mkvurl = mkv
		err := info.Update()
		if err != nil {
			this.showmsg("更新数据失败...")
			return
		} else {
			count, _ := info.Query().Filter("mid", mid).Count()

			minfo.Id = mid
			minfo.Hasepisode = count
			if count == minfo.Episode && len(info.Hdtvurl) > 0 {
				minfo.Isend = 1
			} else {
				minfo.Isend = 0
			}
			minfo.Update("Hasepisode", "Isend")
			this.showmsg("更新数据成功...", this.admindir+"down/add/"+fmt.Sprintf("%d", mid)+"/")
		}
	} else {
		info.Hdtvurl = hdtv
		info.Mkvurl = mkv
		err := info.Insert()
		if err != nil {
			this.showmsg("添加数据失败...")
			return
		} else {
			count, _ := info.Query().Filter("mid", mid).Count()
			minfo.Id = mid
			minfo.Hasepisode = count
			if count == minfo.Episode && len(info.Hdtvurl) > 0 {
				minfo.Isend = 1
			} else {
				minfo.Isend = 0
			}
			minfo.Update("Hasepisode", "Isend")
			this.showmsg("添加数据成功...", this.admindir+"down/add/"+fmt.Sprintf("%d", mid)+"/")
		}
	}

	//this.Ctx.WriteString("save")
}

//list down
func (this *DownaddrHandel) List() {
	this.TplName = "admin/downlist.html"
}
