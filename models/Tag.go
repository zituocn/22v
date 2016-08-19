package models

import (
	"time"
)

//影片标签
///
type TagInfo struct {
	Id      int64
	Mid     int64
	Name    string    `orm:"size(100)"`
	Addtime time.Time `orm:"auto_now_add;type(datetime)"`
}
