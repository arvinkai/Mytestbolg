package controllers

import (
	"github.com/astaxie/beego"
)

type ExitController struct {
	beego.Controller
}

func (this *ExitController) Get() {
	this.Ctx.SetCookie("uname", "", -1, "/")
	this.Ctx.SetCookie("pwd", "", -1, "/")
	this.Redirect("/", 301)
	return
}
