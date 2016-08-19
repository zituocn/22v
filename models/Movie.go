package models

import (
	"fmt"
	"github.com/astaxie/beego/orm"
	"time"
)

//影片模型
type MovieInfo struct {
	Id          int64
	Name        string    `orm:"size(500)"` //影片中文名称
	Ename       string    `orm:"size(100)"` //影片英文名称
	Cid         int64     //分类id
	Cname       string    //分类名称非数据库字段
	Photo       string    `orm:"size(500)"` //海报
	Iphoto      string    `orm:"size(500)"` //推荐到首页的图片
	Actor       string    `orm:"size(500)"` //主演
	Director    string    `orm:"size(500)"` //导演
	Writer      string    `orm:"size(500)"` //编剧
	Language    string    `orm:"size(20)"`  //语言
	Updateweek  int64     //更新星期x
	Playdate    string    `orm:"size(20)"` //开始播放时间
	Content     string    `orm:"size(2000)"`
	Title       string    `orm:"size(500)"`  //seo标题
	Keywords    string    `orm:"size(500)"`  //seo关键字
	Description string    `orm:"size(1000)"` //seo说明
	Views       int64     //浏览量
	Monthviews  int64     //月浏览量
	Status      int64     //影片类型:0为普通 -1为不可见 1为推荐到首页
	Episode     int64     //总集数
	Hasepisode  int64     //已经更新n集
	Isnew       int64     //是否新剧，即新剧推荐
	Addtime     time.Time `orm:"auto_now_add;type(datetime)"` //入库时间
	Updatetime  time.Time `orm:"auto_now_add;type(datetime)"` //更新某季下载的时间，用来排序最近更新
	Editor      string    `orm:"size(50)"`                    //责任编辑
	Isend       int64     //是否已更新完结
}

func (m *MovieInfo) TableName() string {
	return "movie_info"
}

func (m *MovieInfo) Query() orm.QuerySeter {
	return orm.NewOrm().QueryTable(m)
}

func (m *MovieInfo) Insert() error {
	if _, err := orm.NewOrm().Insert(m); err != nil {
		return err
	}
	return nil
}

func (m *MovieInfo) Read(fields ...string) error {
	if err := orm.NewOrm().Read(m, fields...); err != nil {
		return err
	}
	return nil
}

func (m *MovieInfo) Update(fields ...string) error {
	if _, err := orm.NewOrm().Update(m, fields...); err != nil {
		return err
	}
	return nil
}

func (m *MovieInfo) Delete() error {
	if _, err := orm.NewOrm().Delete(m); err != nil {
		return err
	}
	return nil
}

///首页小列表
func (m *MovieInfo) GetList(cid int64, pagesize int64) []*MovieInfo {
	var info MovieInfo
	list := make([]*MovieInfo, 0)
	if cid != 100 && cid != 200 {
		info.Query().OrderBy("-id").Filter("cid", cid).Limit(pagesize, 0).All(&list, "photo", "Id", "Title", "Name", "Addtime", "Hasepisode", "Episode")
	} else if cid == 100 {
		info.Query().OrderBy("-id").Limit(pagesize, 0).All(&list, "photo", "Id", "Title", "Name", "Addtime", "Hasepisode", "Episode")
	} else {
		info.Query().OrderBy("-views").Limit(pagesize, 0).All(&list, "photo", "Id", "Title", "Name", "Addtime", "Hasepisode", "Episode")
	}
	return list
}

//首页今日(星期)更新小列表
func (m *MovieInfo) GetWeekList(pagesize int64) []*MovieInfo {
	var info MovieInfo
	list := make([]*MovieInfo, 0)
	query := info.Query()
	query = query.Filter("isend", 0)
	query = query.Filter("Hasepisode__gte", 1)
	query = query.Filter("status__gte", 0)
	week := time.Now().Weekday()
	if week.String() == "Sunday" {
		week = 6
	}
	//fmt.Println(week)
	query = query.Filter("updateweek", week)
	query.Limit(pagesize, 0).OrderBy("-Id").All(&list)
	return list
}

///内页热门列表
func (m *MovieInfo) GetHotList(cid int64, pagesize int64) []*MovieInfo {
	var (
		info MovieInfo
	)
	list := make([]*MovieInfo, 0)
	if cid > 0 && cid != 100 && cid != 200 {
		info.Query().Filter("status__gte", 0).Filter("cid", cid).OrderBy("-views").Limit(pagesize, 0).All(&list, "Id", "Title", "Name", "Addtime", "photo", "Hasepisode", "episode")
	} else {
		info.Query().Filter("status__gte", 0).OrderBy("-views").Limit(pagesize, 0).All(&list, "Id", "Title", "Name", "Addtime", "photo", "Hasepisode", "episode")
	}
	return list
}

///随机列表
func (m *MovieInfo) GetRandList(pagesize int64) []*MovieInfo {
	list := make([]*MovieInfo, 0)
	o := orm.NewOrm()
	_, _ = o.Raw("SELECT * FROM movie_info AS t1 JOIN (SELECT ROUND(RAND() * ((SELECT MAX(Id) FROM movie_info)-(SELECT MIN(Id) FROM movie_info)) + (SELECT MIN(Id) FROM movie_info)) AS Id) AS t2 WHERE t1.Id >= t2.Id ORDER BY t1.Id LIMIT 6").QueryRows(&list)
	return list
}

//首页的今日更新集数
//模板函数
func GetMovieUpdateEP(hasep int64, ep int64) string {
	var result string
	if hasep < ep {
		result = fmt.Sprintf("%d / %d", hasep+1, ep)
	} else {
		result = fmt.Sprintf("%d / %d", hasep, ep)
	}
	return result
}

//模板函数
//另外一个更新集数
func GetMovieUpdateEPString(hasep int64, ep int64) string {
	var result string
	if hasep == 0 {
		result = "未更新"
		return result
	}
	if hasep == ep {
		result = "已完结"
	} else {
		result = fmt.Sprintf("更新到第%d集", hasep)
	}
	return result
}

//宣传海报
func GetIPhoto(url string, title string) string {
	var result string
	if len(url) > 0 {
		result = "<p><img src=\"" + url + "\" alt=\"" + title + "\" /></p>"
	}
	return result
}
