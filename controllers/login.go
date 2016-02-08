package controllers

import (
	"fmt"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context"
)

type LoginController struct {
	beego.Controller
}

func (this *LoginController) Get() {
	fmt.Println(this.Input())
	IsExit := this.Input().Get("exit") == "true"
	fmt.Println("Login Get:", "Input:", this.Input().Get("exit"), "IsExit:", IsExit)
	if IsExit {
		this.Ctx.SetCookie("uname", "", -1, "/")
		this.Ctx.SetCookie("pwd", "", -1, "/")
		this.Redirect("/", 301)
		return
	}
	this.TplName = "login.html"
	return
}

func (this *LoginController) Post() {
	uname := this.Input().Get("uName")
	pwd := this.Input().Get("pwd")
	auto := this.Input().Get("autoLogin") == "on"
	if beego.AppConfig.String("uName") == uname &&
		beego.AppConfig.String("uPwd") == pwd {
		if auto {
			maxAge := 1<<31 - 1
			this.Ctx.SetCookie("uname", uname, maxAge, "/")
			this.Ctx.SetCookie("pwd", pwd, maxAge, "/")
		} else {
			this.Ctx.SetCookie("uname", uname, 10, "/")
			this.Ctx.SetCookie("pwd", pwd, 10, "/")
		}
	}
	this.Redirect("/", 301)
	return
}

func checkAccount(ctx *context.Context) bool {
	ck, err := ctx.Request.Cookie("uname") //ctx.GetCookie("uname")
	//	uname := ctx.GetCookie("uname")
	if err != nil {
		return false
	}
	uname := ck.Value

	ck, err = ctx.Request.Cookie("pwd")
	//	pwd := ctx.GetCookie("pwd")
	if err != nil {
		return false
	}

	pwd := ck.Value

	if beego.AppConfig.String("uName") == uname &&
		beego.AppConfig.String("uPwd") == pwd {
		return true
	}
	return false
}
