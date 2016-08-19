package models

import (
	"github.com/astaxie/beego/orm"
	"strings"
	"time"
)

///前台推荐控制
type RecommendInfo struct {
	Id      int64
	Name    string    `orm:"size(200)"`
	Content string    `orm:"size(8000)"`
	Addtime time.Time `orm:"auto_now_add;type(datetime)"`
}

func (m *RecommendInfo) TableName() string {
	return "recommend_info"
}

func (m *RecommendInfo) Query() orm.QuerySeter {
	return orm.NewOrm().QueryTable(m)
}

func (m *RecommendInfo) Insert() error {
	if _, err := orm.NewOrm().Insert(m); err != nil {
		return err
	}
	return nil
}

func (m *RecommendInfo) Read(fields ...string) error {
	if err := orm.NewOrm().Read(m, fields...); err != nil {
		return err
	}
	return nil
}

func (m *RecommendInfo) Update(fields ...string) error {
	if _, err := orm.NewOrm().Update(m, fields...); err != nil {
		return err
	}
	return nil
}

func (m *RecommendInfo) Delete() error {
	if _, err := orm.NewOrm().Delete(m); err != nil {
		return err
	}
	return nil
}

func (m *RecommendInfo) GetContent(id int64) string {
	var (
		info RecommendInfo
		err  error
	)
	var result string
	info.Id = id
	err = info.Read()
	if err != nil {
		result = ""
	}
	result = strings.Replace(info.Content, "<p>", "", -1)
	result = strings.Replace(result, "</p>", "", -1)
	return result
}
