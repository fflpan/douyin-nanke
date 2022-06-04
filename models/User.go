package models

type User struct {
	Id            int64  `json:"id" orm:"pk;auto;column(user_id);size(10)"`
	Name          string `json:"name" orm:"column(username);size(255)"`
	Pass          string `json:"-" orm:"column(password);size(255)"`
	FollowCount   int64  `json:"follow_count" orm:"column(follow_count);size(10)"`
	FollowerCount int64  `json:"follower_count" orm:"column(follower_count);size(10)"`
	IsFollow      bool   `json:"is_follow" orm:"-"`
}
