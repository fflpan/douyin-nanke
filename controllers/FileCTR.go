package controllers

import (
	"DouYin/models"
	"fmt"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	"os"
	"os/exec"
	"strconv"
	"strings"
	"time"
)

const ADDRESS = "http://10.180.139.161:8080/"
const FF = "/Users/mac/go/tools/ffmpeg"

type FileCTR struct {
	beego.Controller
}

var UpMap = make(map[string]interface{}, 2)

func (ctr FileCTR) Upload() {
	token := ctr.Ctx.Input.Query("token")
	title := ctr.Ctx.Input.Query("title")
	_, header, _ := ctr.GetFile("data")
	filename := header.Filename
	split := strings.Split(filename, "/")
	if len(split) > 1 {
		filename = split[len(split)-1]
	}
	ss := strconv.FormatInt(time.Now().UnixNano(), 10)
	filename = ss + filename
	pigName := ss + ".jpg"
	ctr.SaveToFile("data", "static/video/"+filename)
	c := exec.Command(FF, "-i", "static/video/"+filename,
		"-ss", "1", "-f", "image2", "'-frames:v 1'", "static/img/"+pigName)
	c.Stderr = os.Stderr
	c.Run()
	UpMap["status_code"] = "200"
	UpMap["status_msg"] = "success"

	//上传成功 将视频信息写入数据表
	go upLoad(filename, pigName, token, title)
	ctr.Data["json"] = UpMap
	ctr.ServeJSON()
}

func upLoad(filename, pigName, token, title string) {
	filePath := ADDRESS + "video/" + filename
	coverPath := ADDRESS + "img/" + pigName
	o := orm.NewOrm()
	exec, err := o.Raw("insert into video(author_id,play_url,cover_url,"+
		"pub_time,title) values(?,?,?,now(),?)",
		GetUserIdByToken(token), filePath, coverPath, title).Exec()
	if err != nil {
		fmt.Println(err, exec)
	}
}

var ListMap = make(map[string]interface{}, 3)

func (ctr FileCTR) List() {
	token := ctr.Ctx.Input.Query("token")
	ListMap["status_code"] = "0"
	ListMap["status_msg"] = "success"
	o := orm.NewOrm()
	var videos []models.Video
	_, err := o.Raw("select * from video where author_id=? order by video_id desc",
		GetUserIdByToken(token)).QueryRows(&videos)
	if err != nil {
		fmt.Errorf("%s", err)
	}
	userId := GetUserIdByToken(token)
	for i := 0; i < len(videos); i++ {
		videos[i].IsFavorite = isFavorite(userId, videos[i].Id)
		videos[i].Author = QueryUserById(videos[i].AuthorId, token)
	}
	ListMap["video_list"] = videos
	ctr.Data["json"] = ListMap
	ctr.ServeJSON()
}

var FeedMap = make(map[string]interface{}, 3)

func (ctr FileCTR) Feed() {
	token := ctr.Ctx.Input.Query("token")
	FeedMap["status_code"] = "0"
	FeedMap["status_msg"] = "success"
	o := orm.NewOrm()
	var videos []models.Video
	latestTime := ctr.Ctx.Input.Query("latest_time")
	if latestTime == "" {
		_, err := o.Raw("select * from video limit 30").QueryRows(&videos)
		if err != nil {
			fmt.Errorf("%s", err)
		}
	} else {
		parse, _ := time.Parse("2006-01-02 15-04-05", latestTime)
		_, err := o.Raw("select * from video where pub_time>=? order by video_id desc limit 30 ",
			parse).QueryRows(&videos)
		if err != nil {
			fmt.Errorf("%s", err)
		}
	}
	userId := GetUserIdByToken(token)
	for i := 0; i < len(videos); i++ {
		videos[i].IsFavorite = isFavorite(userId, videos[i].Id)
		videos[i].Author = QueryUserById(videos[i].AuthorId, token)
	}
	FeedMap["video_list"] = videos
	ctr.Data["json"] = FeedMap
	ctr.ServeJSON()
}

func isFavorite(userId, videoId int64) bool {
	o := orm.NewOrm()
	var favorite = models.Favorite{}
	o.Raw("Select * from favorite where user_id=? and video_id=? limit 1",
		userId, videoId).QueryRow(&favorite)
	return favorite.Id != 0
}

var ResMap = make(map[string]interface{}, 2)

func (ctr FileCTR) Favorite() {
	//userId := ctr.Ctx.Input.Query("user_id")
	token := ctr.Ctx.Input.Query("token")
	videoId := ctr.Ctx.Input.Query("video_id")
	actionType := ctr.Ctx.Input.Query("action_type")
	o := orm.NewOrm()
	UId := GetUserIdByToken(token)
	VId, _ := strconv.Atoi(videoId)
	if actionType == "1" {
		o.Raw("Insert into favorite(user_id,video_id) values(?,?)",
			UId, int64(VId)).Exec()
	} else {
		o.Raw("delete from favorite where user_id=? and video_id=?",
			UId, int64(VId)).Exec()
	}
	go favorite(VId, actionType)
	ResMap["status_code"] = 0
	ResMap["status_msg"] = "success"
	ctr.Data["json"] = ResMap
	ctr.ServeJSON()
}

func favorite(VId int, action string) {
	o := orm.NewOrm()
	var ac int
	if action == "1" {
		ac = 1
	} else {
		ac = -1
	}
	o.Raw("update video set favorite_count=favorite_count+? where video_id=?",
		ac, int64(VId)).Exec()
}

var FavMap = make(map[string]interface{}, 2)

func (ctr FileCTR) FavList() {
	//userId := ctr.Ctx.Input.Query("user_id")
	token := ctr.Ctx.Input.Query("token")
	UId := GetUserIdByToken(token)
	o := orm.NewOrm()
	var videos []models.Video
	o.Raw("select * from video where video_id in (select video_id from favorite"+
		" where user_id=?) order by video_id desc", UId).QueryRows(&videos)
	FavMap["status_code"] = "0"
	FavMap["status_msg"] = "success"
	for i := 0; i < len(videos); i++ {
		videos[i].IsFavorite = isFavorite(UId, videos[i].Id)
		videos[i].Author = QueryUserById(videos[i].AuthorId, token)
	}
	FavMap["video_list"] = videos
	ctr.Data["json"] = FavMap
	ctr.ServeJSON()
}
