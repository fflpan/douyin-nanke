package controllers

import "github.com/astaxie/beego"

type ErrorCTR struct {
	beego.Controller
}

var Unk = make(map[string]interface{}, 2)

func (ctr ErrorCTR) UnknownToken() {
	Unk["status_code"] = "400"
	Unk["status_msg"] = "unknown tokens \n please login in first"
	ctr.Data["json"] = Unk
	ctr.ServeJSON()
}
