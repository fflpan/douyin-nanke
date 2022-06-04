package models

type Follow struct {
	Id           int64 `orm:"pk;auto;column(follow_id);size(10)"`
	FollowerId   int64 `orm:"column(follower_id);size(10)"`
	BeFollowedId int64 `orm:"column(be_follow_id);size(10)"`
}
