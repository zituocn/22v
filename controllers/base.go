package controllers

import (
	//"fmt"
	"github.com/astaxie/beego"
	"github.com/zituocn/VMovie/models"
	"strconv"
	"strings"
	"time"
)

type baseController struct {
	beego.Controller
}

func (this *baseController) Prepare() {
	var (
		subnav      string
		hotkeywords string
		rinfo       models.RecommendInfo
	)
	subnav = rinfo.GetContent(2)
	hotkeywords = rinfo.GetContent(3)
	this.Data["subnav"] = subnav
	this.Data["hotkeywords"] = hotkeywords

	host := "https://m.22v.net"
	ug := this.Ctx.Request.UserAgent()
	path := this.Ctx.Request.URL.String()

	if !strings.Contains(path, "api") {
		ug = strings.ToLower(ug)
		if strings.Contains(ug, "iphone") || strings.Contains(ug, "android") || strings.Contains(ug, "phone") {
			this.Ctx.Redirect(302, host+path)
		}
	}
}

func Error(err error) {
	if err != nil {
		panic(err)
		beego.Error(err.Error())
		//os.Exit(1)
	}
}

//返回星期四格式
func (this *baseController) GetWeekString() string {
	var wstring string
	switch time.Now().Weekday() {
	case time.Monday:
		wstring = "一"
	case time.Tuesday:
		wstring = "二"
	case time.Wednesday:
		wstring = "三"
	case time.Thursday:
		wstring = "四"
	case time.Friday:
		wstring = "五"
	case time.Saturday:
		wstring = "六"
	case time.Sunday:
		wstring = "天"
	default:
		wstring = "X"
	}
	return wstring
}

//显示分页链接
func (this *baseController) PageList(pagesize, page, recordcount int64, first bool, path string) (pager string) {
	if recordcount == 0 {
		return ""
	}

	var pagecount int64
	pagecount = 0

	if recordcount%pagesize == 0 {
		pagecount = recordcount / pagesize
	} else {
		pagecount = (recordcount / pagesize) + 1
	}

	pager = "<span>" + strconv.FormatInt(page, 10) + "/" + strconv.FormatInt(pagecount, 10) + "</span>"

	if pagecount < 2 {
		return "<span>共1页</span>"
	}

	pager = pager + "<a href=\"" + path + "/\">第一页</a>"

	if page > 1 {
		if page == 2 {
			pager = pager + "<a href=\"" + path + "/\">上一页</a>"
		} else {
			pager = pager + "<a href=\"" + path + "/" + strconv.FormatInt(page-1, 10) + "/\" >上一页</a>"
		}
	} else {
		pager = pager + "<a href=\"" + path + "/\">上一页</a>"
	}

	if page < pagecount {
		pager = pager + "<a href=\"" + path + "/" + strconv.FormatInt(page+1, 10) + "/\" class=\"next\">下一页</a>"
	} else {
		pager = pager + "<a href=\"" + path + "/" + strconv.FormatInt(pagecount, 10) + "/\"  class=\"next\">下一页</a>"
	}

	pager = pager + "<a href=\"" + path + "/" + strconv.FormatInt(pagecount, 10) + "/\"  class=\"next\">最后一页</a>"

	pager = pager + "<span >每页" + strconv.FormatInt(pagesize, 10) + "/共" + strconv.FormatInt(recordcount, 10) + "</span>"

	return pager

}
