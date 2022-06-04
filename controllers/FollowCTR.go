package controllers

import (
	"DouYin/models"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	"strconv"
)

type FollowCTR struct {
	beego.Controller
}

var LikeMap = make(map[string]interface{}, 2)

func (ctr FollowCTR) Like() {
	//UID := ctr.Ctx.Input.Query("user_id")
	token := ctr.Ctx.Input.Query("token")
	UID2 := ctr.Ctx.Input.Query("to_user_id")
	action := ctr.Ctx.Input.Query("action_type")

	userId := GetUserIdByToken(token)
	toUserId, _ := strconv.Atoi(UID2)
	act, _ := strconv.Atoi(action)

	o := orm.NewOrm()
	if act == 1 {
		o.Raw("insert into follow(follower_id,be_follow_id) values(?,?)", userId, toUserId).Exec()
		o.Raw("update user set follow_count=follow_count+1 where user_id=?", userId).Exec()
		o.Raw("update user set follower_count=follower_count+1 where user_id=?", toUserId).Exec()
		models.RO.Client.HIncrBy(models.RO.Ctx, token, "follow_count", 1)
	} else {
		o.Raw("delete from follow where follower_id=? and be_follow_id=?", userId, toUserId).Exec()
		o.Raw("update user set follow_count=follow_count-1 where user_id=?", userId).Exec()
		o.Raw("update user set follower_count=follower_count-1 where user_id=?", toUserId).Exec()
		models.RO.Client.HIncrBy(models.RO.Ctx, token, "follow_count", -1)
	}
	LikeMap["status_code"] = "0"
	LikeMap["status_msg"] = "success"
	ctr.Data["json"] = LikeMap
	ctr.ServeJSON()
}

var LikeListMap = make(map[string]interface{}, 3)

func (ctr FollowCTR) LikeList() {
	userId := ctr.Ctx.Input.Query("user_id")
	o := orm.NewOrm()
	var users []models.User
	o.Raw("select * from user where user_id in (select be_follow_id from follow where follower_id = ?)",
		userId).QueryRows(&users)
	for i := 0; i < len(users); i++ {
		users[i].IsFollow = true
	}
	LikeListMap["status_code"] = "0"
	LikeListMap["status_msg"] = "success"
	LikeListMap["user_list"] = users
	ctr.Data["json"] = LikeListMap
	ctr.ServeJSON()
}

func (ctr FollowCTR) BeLikedList() {
	userId := ctr.Ctx.Input.Query("user_id")
	token := ctr.Ctx.Input.Query("token")
	var users []models.User
	o := orm.NewOrm()
	o.Raw("select * from user where user_id in (select follower_id from follow where be_follow_id = ?)",
		userId).QueryRows(&users)
	for i := 0; i < len(users); i++ {
		users[i].IsFollow = IsFollow(GetUserIdByToken(token), users[i].Id)
	}
	LikeListMap["status_code"] = "0"
	LikeListMap["status_msg"] = "success"
	LikeListMap["user_list"] = users
	ctr.Data["json"] = LikeListMap
	ctr.ServeJSON()
}
