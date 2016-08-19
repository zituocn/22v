package models

import (
	"github.com/astaxie/beego/orm"
	"time"
)

//页面/专题模型
//手动模式 自定义程度较高的专属页面
type PageInfo struct {
	Id          int64
	Name        string    `orm:"size(200)"`   //中文专题名称
	Ename       string    `orm:"sisze(200)"`  //英语说明  /page/newsyear/
	Content     string    `orm:"size(10000)"` //文章正文 html
	Title       string    `orm:"size(300)"`   //seo标题
	Description string    `orm:"size(500)"`   //seo页面说明
	Keywords    string    `orm:"size(200)"`   //seo关键字
	Status      int64     //状态 -1时，前台不显示 0 为正常
	Views       int64     //浏览量
	Editor      string    `orm:"size(20)"` //责任编辑 显示在页面上
	Addtime     time.Time `orm:"auto_now_add;type(datetime)"`
}

func (m *PageInfo) TableName() string {
	return "page_info"
}

func (m *PageInfo) Query() orm.QuerySeter {
	return orm.NewOrm().QueryTable(m)
}

func (m *PageInfo) Insert() error {
	if _, err := orm.NewOrm().Insert(m); err != nil {
		return err
	}
	return nil
}

func (m *PageInfo) Read(fields ...string) error {
	if err := orm.NewOrm().Read(m, fields...); err != nil {
		return err
	}
	return nil
}

func (m *PageInfo) Update(fields ...string) error {
	if _, err := orm.NewOrm().Update(m, fields...); err != nil {
		return err
	}
	return nil
}

func (m *PageInfo) Delete() error {
	if _, err := orm.NewOrm().Delete(m); err != nil {
		return err
	}
	return nil
}

//最近N条资讯
func (m *PageInfo) GetTopNews(pagesize int64) []*PageInfo {
	var (
		info PageInfo
	)
	list := make([]*PageInfo, 0)
	info.Query().OrderBy("-Id").Filter("status", 0).Limit(pagesize, 0).All(&list, "Id", "Title", "Name", "Addtime")
	return list
}
