package controllers

import (
	"DouYin/models"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	"strconv"
)

type UserCTR struct {
	beego.Controller
}

var ReMap = make(map[string]string, 5)

func (ctr UserCTR) Register() {
	username := ctr.Ctx.Input.Query("username")
	password := ctr.Ctx.Input.Query("password")
	var o = orm.NewOrm()
	user := models.User{}
	user.Name = username
	//检查用户是否已经存在
	o.Raw("select * from user where username=?", username).QueryRow(&user)
	if user.Id == 0 {
		// 如不存在
		user.Pass = password
		o.Insert(&user)
		ReMap["status_code"] = "0"
		ReMap["status_msg"] = "success"
		ReMap["user_id"] = strconv.Itoa(int(user.Id))
		ReMap["token"] = username
		// 将其加入tokens
		addUser(&user)
	} else {
		//如已经存在
		ReMap["status_code"] = "1"
		ReMap["status_msg"] = "username already exist \n please make a change"
		ReMap["user_id"] = "0"
		ReMap["token"] = "0"
	}
	ctr.Data["json"] = ReMap
	ctr.ServeJSON()
}

var LoginMap = make(map[string]string, 5)

func (ctr UserCTR) Login() {
	username := ctr.Ctx.Input.Query("username")
	password := ctr.Ctx.Input.Query("password")
	var o = orm.NewOrm()
	user := models.User{}
	o.Raw("select * from user where username=? and password=? limit 1",
		username, password).QueryRow(&user)
	if user.Name != username {
		//如不存在
		LoginMap["status_code"] = "1"
		LoginMap["status_msg"] = "username/password not match \n please check"
		LoginMap["user_id"] = "0"
		LoginMap["token"] = "0"
	} else {
		//如存在
		LoginMap["status_code"] = "0"
		LoginMap["status_msg"] = "success"
		LoginMap["user_id"] = strconv.Itoa(int(user.Id))
		LoginMap["token"] = username
		// 将其加入tokens
		addUser(&user)
	}

	ctr.Data["json"] = LoginMap
	ctr.ServeJSON()
}

var QueryMaps = make(map[string]interface{}, 3)

func (ctr UserCTR) QueryById() {
	userId := ctr.Ctx.Input.Query("user_id")
	token := ctr.Ctx.Input.Query("token")
	AToI, _ := strconv.Atoi(userId)
	user := QueryUserById(int64(AToI), token)
	QueryMaps["status_code"] = "0"
	QueryMaps["status_msg"] = "success"
	QueryMaps["user"] = user
	ctr.Data["json"] = QueryMaps
	ctr.ServeJSON()
}

func addUser(user *models.User) {
	models.RO.Client.HSet(models.RO.Ctx, user.Name, "id", user.Id)
	models.RO.Client.HSet(models.RO.Ctx, user.Name, "user_name", user.Name)
	models.RO.Client.HSet(models.RO.Ctx, user.Name, "follow_count", user.FollowerCount)
	models.RO.Client.HSet(models.RO.Ctx, user.Name, "follower_count", user.FollowerCount)
}

func IsFollow(follower, beFollower int64) bool {
	o := orm.NewOrm()
	var follow = models.Follow{}
	o.Raw("Select * from follow where follower_id=? and be_follow_id=? limit 1",
		follower, beFollower).QueryRow(&follow)
	return follow.Id != 0
}

func GetUserIdByToken(token string) int64 {
	get, _ := models.RO.Client.HGet(models.RO.Ctx, token, "id").Result()
	userId, _ := strconv.Atoi(get)
	return int64(userId)
}

func QueryUserById(userId int64, token string) *models.User {
	o := orm.NewOrm()
	user := models.User{}
	o.Raw("select * from user where user_id=? limit 1",
		userId).QueryRow(&user)
	user.IsFollow = IsFollow(GetUserIdByToken(token), userId)
	return &user
}
