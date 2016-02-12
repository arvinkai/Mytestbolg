package controllers

import (
	"Myblog/models"

	"github.com/astaxie/beego"
)

type TopicController struct {
	beego.Controller
}

func (this *TopicController) Get() {
	this.Data["IsLogin"] = checkAccount(this.Ctx)
	this.Data["IsTopic"] = true
	this.TplName = "topic.html"
	topics, err := models.GetAllTopics(true)
	if err != nil {
		beego.Error(err.Error())
	} else {
		this.Data["Topics"] = topics
	}
}

func (this *TopicController) Add() {
	this.TplName = "topic_add.html"
}

func (this *TopicController) Post() {
	if !checkAccount(this.Ctx) {
		this.Redirect("/login", 302)
		return
	}

	title := this.Input().Get("title")
	content := this.Input().Get("content")

	err := models.AddTopic(title, content)
	if err != nil {
		return
	}

	this.Redirect("/topic", 302)
}
