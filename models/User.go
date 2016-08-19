package models

import (
	"github.com/astaxie/beego/orm"
	"time"
)

type UserInfo struct {
	Id            int64
	Username      string    `orm:"size(20)"`
	Password      string    `orm:"size(32)"`
	Nickname      string    `orm:"size(40)"`
	Lastlogintime time.Time `orm:"auto_now_add;type(datetime)"`
	Logintimes    int64
	Lastloginip   string    `orm:"size(32)"`
	Addtime       time.Time `orm:"auto_now_add;type(datetime)"`
}

func (m *UserInfo) TableName() string {
	return "user_info"
}

func (m *UserInfo) Insert() error {
	if _, err := orm.NewOrm().Insert(m); err != nil {
		return err
	}
	return nil
}

func (m *UserInfo) Read(fields ...string) error {
	if err := orm.NewOrm().Read(m, fields...); err != nil {
		return err
	}
	return nil
}

func (m *UserInfo) Update(fields ...string) error {
	if _, err := orm.NewOrm().Update(m, fields...); err != nil {
		return err
	}
	return nil
}

func (m *UserInfo) Delete() error {
	if _, err := orm.NewOrm().Delete(m); err != nil {
		return err
	}
	return nil
}

func (m *UserInfo) Query() orm.QuerySeter {
	return orm.NewOrm().QueryTable(m)
}
