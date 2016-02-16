package controllers

import (
	"Myblog/models"

	"github.com/astaxie/beego"
)

type ReplyController struct {
	beego.Controller
}

func (this *ReplyController) Add() {
	tId := this.Input().Get("tid")
	nickName := this.Input().Get("nickname")
	reContent := this.Input().Get("content")

	if tId != "" {
		err := models.AddReply(tId, nickName, reContent)
		if err == nil {
			dir := "/topic/view/" + tId
			this.Redirect(dir, 302)
			return
		}
	}

	this.Redirect("/", 302)
}

func (this *ReplyController) Delete() {
	Id := this.Input().Get("id")
	tId := this.Input().Get("tid")
	if Id != "" {
		err := models.DelReply(Id)
		if err == nil {
			dir := "/topic/view/" + tId
			this.Redirect(dir, 302)
			return
		}
	}

	this.Redirect("/", 302)
}
