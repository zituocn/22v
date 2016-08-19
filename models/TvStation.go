package models

//电视台信息
type TvStationInfo struct {
	Id          int64
	Name        string `orm:"size(50)"`
	Ename       string `orm:"size(50)"`
	Title       string `orm:"size(300)"`
	Description string `orm:"size(300)"`
}
