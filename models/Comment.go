package models

import (
	"time"
)

type Comment struct {
	Cid        int64     `json:"id" orm:"auto;pk;column(cid);size(10)"`
	UserId     int64     `json:"-" orm:"column(user_id);size(10)"`
	User       *User     `json:"user" orm:"-"`
	VideoId    int64     `json:"-" orm:"column(video_id);size(10)"`
	Content    string    `json:"content" orm:"column(content);size(10)"`
	Date       time.Time `json:"-" orm:"column(create_date);size(10)"`
	CreateDate string    `json:"create_date" orm:"-"`
}
