package models

import (
	"github.com/astaxie/beego/orm"
	"time"
)

///影片关系数据结构
type RelationInfo struct {
	Id      int64
	Mids    string    `orm:"size(300)"`                   //影片的mids
	Name    string    `orm:"size(50)"`                    //关系名称
	Addtime time.Time `orm:"auto_now_add;type(datetime)"` //关系生成时间
}

func (m *RelationInfo) TableName() string {
	return "relation_info"
}

func (m *RelationInfo) Query() orm.QuerySeter {
	return orm.NewOrm().QueryTable(m)
}

func (m *RelationInfo) Insert() error {
	if _, err := orm.NewOrm().Insert(m); err != nil {
		return err
	}
	return nil
}

func (m *RelationInfo) Read(fields ...string) error {
	if err := orm.NewOrm().Read(m, fields...); err != nil {
		return err
	}
	return nil
}

func (m *RelationInfo) Update(fields ...string) error {
	if _, err := orm.NewOrm().Update(m, fields...); err != nil {
		return err
	}
	return nil
}

func (m *RelationInfo) Delete() error {
	if _, err := orm.NewOrm().Delete(m); err != nil {
		return err
	}
	return nil
}
