package models

type Favorite struct {
	Id      int64 `orm:"auto;pk;column(favorite_id);size(10)"`
	UserId  int64 `orm:"column(user_id);size(10)"`
	VideoId int64 `orm:"column(video_id);size(10)"`
}
