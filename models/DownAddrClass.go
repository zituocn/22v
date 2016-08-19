package models

//下载片源分类
//mp4 720p 中文字幕
//mkv 1080p 中文字幕
//mp4 720 无字幕
type DownAddrClassInfo struct {
	Id    int64
	Name  string `orm:"size(30)"`
	Ename string `orm:"size(30)"`
}
