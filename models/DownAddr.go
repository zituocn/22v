package models

import (
	"github.com/astaxie/beego/orm"
	"time"
)

//下载地址
//@Did 为片源类型 mp4/mkv/中文/无字幕...
type DownAddrInfo struct {
	Id      int64     `json:"id"`
	Mid     int64     `json:"mid"`
	Name    string    `orm:"size(20)" json:"name"`
	Hdtvurl string    `orm:"size(500)" json:"hdtvurl"`
	Mkvurl  string    `orm:"size(500)" json:"mkvurl"`
	Ep      int64     `json:"ep"`
	Addtime time.Time `orm:"auto_now_add;type(datetime)" json:"addtime"`
}

func (m *DownAddrInfo) TableName() string {
	return "down_addr_info"
}

func (m *DownAddrInfo) Query() orm.QuerySeter {
	return orm.NewOrm().QueryTable(m)
}

func (m *DownAddrInfo) Insert() error {
	if _, err := orm.NewOrm().Insert(m); err != nil {
		return err
	}
	return nil
}

func (m *DownAddrInfo) Read(fields ...string) error {
	if err := orm.NewOrm().Read(m, fields...); err != nil {
		return err
	}
	return nil
}

func (m *DownAddrInfo) Update(fields ...string) error {
	if _, err := orm.NewOrm().Update(m, fields...); err != nil {
		return err
	}
	return nil
}

func (m *DownAddrInfo) Delete() error {
	if _, err := orm.NewOrm().Delete(m); err != nil {
		return err
	}
	return nil
}
