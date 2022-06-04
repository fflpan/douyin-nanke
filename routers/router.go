package routers

import (
	"DouYin/controllers"
	"DouYin/models"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context"
)

var WhiteRouter = make(map[string]int, 5)

func init() {

	//路由白名单
	WhiteRouter["/douyin/user/register/"] = 1
	WhiteRouter["/douyin/user/register"] = 1
	WhiteRouter["/douyin/user/login/"] = 1
	WhiteRouter["/douyin/user/login"] = 1
	WhiteRouter["/douyin/feed"] = 1
	WhiteRouter["/douyin/feed/"] = 1
	WhiteRouter["/douyin/comment/list/"] = 1
	WhiteRouter["/douyin/relation/follow/list/"] = 1
	WhiteRouter["/douyin/relation/follower/list/"] = 1

	//路由
	beego.Router("/douyin/user/register", &controllers.UserCTR{},
		"post:Register")
	beego.Router("/douyin/user/login", &controllers.UserCTR{},
		"post:Login")
	beego.Router("/douyin/user", &controllers.UserCTR{},
		"get:QueryById")

	beego.Router("/douyin/publish/action", &controllers.FileCTR{},
		"post:Upload")
	beego.Router("/douyin/publish/list", &controllers.FileCTR{},
		"get:List")
	beego.Router("/douyin/favorite/list", &controllers.FileCTR{},
		"get:FavList")
	beego.Router("/douyin/feed", &controllers.FileCTR{},
		"get:Feed")
	beego.Router("/douyin/favorite/action", &controllers.FileCTR{},
		"post:Favorite")

	beego.Router("/douyin/relation/action/", &controllers.FollowCTR{},
		"post:Like")
	beego.Router("/douyin/relation/follow/list/", &controllers.FollowCTR{},
		"get:LikeList")
	beego.Router("/douyin/relation/follower/list/", &controllers.FollowCTR{},
		"get:BeLikedList")

	beego.Router("/douyin/comment/list/", &controllers.CommentCTR{},
		"get:GetComments")
	beego.Router("/douyin/comment/action/", &controllers.CommentCTR{},
		"post:AddComment")

	beego.Router("/unknownToken", &controllers.ErrorCTR{},
		"*:UnknownToken")

	//过滤器 鉴定token
	beego.InsertFilter("/*", beego.BeforeRouter, TokenFilter)

}

func TokenFilter(ctx *context.Context) {
	if WhiteRouter[ctx.Request.URL.Path] != 1 {
		token := ctx.Input.Query("token")
		result, _ := models.RO.Client.Exists(models.RO.Ctx, token).Result()
		if result != 1 {
			ctx.Request.URL.Path = "/unknownToken"
		}
	}
}
