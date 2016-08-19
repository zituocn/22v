package models

import (
	//"fmt"
	"github.com/astaxie/beego/orm"
)

//影片类型
//动作/科幻/情景/罪案/惊悚
type MovieClassInfo struct {
	Id          int64
	Name        string `orm:"size(50)"`
	Ename       string `orm:"size(50)"`
	Title       string `orm:"size(200)"`
	Keywords    string `orm:"size(200)"`
	Description string `orm:"size(300)"`
}

func (m *MovieClassInfo) TableName() string {
	return "movie_class_info"
}

func (m *MovieClassInfo) Query() orm.QuerySeter {
	return orm.NewOrm().QueryTable(m)
}

func (m *MovieClassInfo) Insert() error {
	if _, err := orm.NewOrm().Insert(m); err != nil {
		return err
	}
	return nil
}

func (m *MovieClassInfo) Read(fields ...string) error {
	if err := orm.NewOrm().Read(m, fields...); err != nil {
		return err
	}
	return nil
}

func (m *MovieClassInfo) Update(fields ...string) error {
	if _, err := orm.NewOrm().Update(m, fields...); err != nil {
		return err
	}
	return nil
}

func (m *MovieClassInfo) Delete() error {
	if _, err := orm.NewOrm().Delete(m); err != nil {
		return err
	}
	return nil
}

func GetCacheList() []*MovieClassInfo {
	var info MovieClassInfo
	key := "classlist"
	list := make([]*MovieClassInfo, 0)
	err := GetCache(key, &list)
	if err != nil {
		info.Query().All(&list)
		SetCache(key, list)
	}
	return list
}

func GetMovieClassInfoByCid(cid int64) *MovieClassInfo {
	list := GetCacheList()
	info := new(MovieClassInfo)
	for _, v := range list {
		if v.Id == cid {
			info = v
			return info
		}
	}
	return info
}

func GetMovieClassNameByCid(cid int64) string {
	info := GetMovieClassInfoByCid(cid)
	if info != nil && info.Id > 0 {
		return info.Name
	}
	return ""
}
