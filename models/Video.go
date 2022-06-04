package models

import "time"

type Video struct {
	Id            int64     `json:"id" orm:"pk;column(video_id);auto;size(10)"`
	AuthorId      int64     `json:"-" orm:"column(author_id);size(10)"`
	Author        *User     `json:"author" orm:"-"`
	PlayUrl       string    `json:"play_url,omitempty" orm:"column(play_url);size(255)"`
	CoverUrl      string    `json:"cover_url,omitempty" orm:"column(cover_url);size(255)"`
	FavoriteCount int64     `json:"favorite_count" orm:"column(favorite_count);size(10)"`
	CommentCount  int64     `json:"comment_count" orm:"column(comment_count);size(10)"`
	IsFavorite    bool      `json:"is_favorite" orm:"-"`
	Time          time.Time `json:"-" orm:"column(pub_time)"`
	Title         string    `json:"title" orm:"column(title);size(300)"`
}
