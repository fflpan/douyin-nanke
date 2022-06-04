package controllers

import (
	"DouYin/models"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	"time"
)

type CommentCTR struct {
	beego.Controller
}

var CommentMap = make(map[string]interface{}, 3)

func (ctr CommentCTR) GetComments() {
	videoId := ctr.Ctx.Input.Query("video_id")
	token := ctr.Ctx.Input.Query("token")
	o := orm.NewOrm()
	var comments []models.Comment
	o.Raw("select * from comment where video_id=?", videoId).QueryRows(&comments)
	for i := 0; i < len(comments); i++ {
		comments[i].User = QueryUserById(comments[i].UserId, token)
		comments[i].CreateDate = comments[i].Date.Format("01-02")
	}
	CommentMap["status_code"] = "0"
	CommentMap["status_msg"] = "success"
	CommentMap["comment_list"] = comments
	ctr.Data["json"] = CommentMap
	ctr.ServeJSON()
}

var ComMap = make(map[string]interface{}, 3)

func (ctr CommentCTR) AddComment() {
	//userId := ctr.Ctx.Input.Query("user_id")
	videoId := ctr.Ctx.Input.Query("video_id")
	action := ctr.Ctx.Input.Query("action_type")
	token := ctr.Ctx.Input.Query("token")
	ComMap["status_code"] = 0
	ComMap["status_msg"] = "success"
	userId := GetUserIdByToken(token)
	o := orm.NewOrm()
	if action == "1" {
		content := ctr.Ctx.Input.Query("comment_text")
		exec, _ := o.Raw("insert into comment(user_id,video_id,content,create_date) "+
			"values(?,?,?,now())", userId, videoId, content).Exec()
		id, _ := exec.LastInsertId()
		o.Raw("update video set comment_count=comment_count+1 where video_id=?", videoId).Exec()
		ComMap["comment"] = &models.Comment{
			Cid:        id,
			Content:    content,
			CreateDate: time.Now().Format("01-02"),
			User:       QueryUserById(userId, token),
		}
	} else {
		CId := ctr.Ctx.Input.Query("comment_id")
		o.Raw("delete from comment where comment_id=?", CId).Exec()
		o.Raw("update video set comment_count=comment_count-1 where video_id=?", videoId).Exec()
		ComMap["comment"] = nil
	}
	ctr.Data["json"] = ComMap
	ctr.ServeJSON()
}
