package main

import (
	"DouYin/models"
	_ "DouYin/routers"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
)

func main() {
	// 连接到本地 mysql 数据库
	orm.RegisterDataBase("default", "mysql",
		"root:fflpan920@tcp(127.0.0.1:3306)/golearn?charset=utf8", 30)

	// 注册模型
	orm.RegisterModel(new(models.User), new(models.Video),
		new(models.Follow), new(models.Favorite), new(models.Comment))

	// 创建表
	orm.RunSyncdb("default", false, true)

	// 连接云服务器redis
	// 连接成功redis中会有一个token
	models.RO.Client.Set(models.RO.Ctx, "token", "fflpan", 600)
	defer models.RO.Client.Close()

	//配置静态文件路径
	beego.SetStaticPath("/video", "static/video")
	beego.SetStaticPath("/img", "static/img")

	beego.Run()
}
