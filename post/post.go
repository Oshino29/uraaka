package post

import (
	"html/template"
)

// Time 暂停为 string，如果以后需要实现时间相关的 feature，再来更改为 time.time，并把 storage/storage.go 中数据库的 time 也改成 sql 中的 time
type Post struct {
	Text string
	Time string
	TextHTML template.HTML
	// 这个 TimeHtml 虽然在其他地方完全没有用到，但是出于不明原因，不能删除
	TimeHTML string
}
type Posts []Post