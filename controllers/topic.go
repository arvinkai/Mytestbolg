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

	tId := this.Input().Get("tid")
	title := this.Input().Get("title")
	content := this.Input().Get("content")

	if "" != tId {
		err := models.ModifyTopic(tId, title, content)
		if err != nil {
			beego.Error(err)
			this.Redirect("/", 302)
			return
		}
	} else {
		err := models.AddTopic(title, content)
		if err != nil {
			beego.Error(err)
			this.Redirect("/", 302)
			return
		}
	}
	this.Redirect("/topic", 302)
}

func (this *TopicController) View() {

	this.TplName = "topic_view.html"
	tId := this.Ctx.Input.Param("0")
	topic, err := models.GetTopic(tId)
	//	fmt.Println(topic)
	//	return
	if err != nil {
		beego.Error(err)
		this.Redirect("/", 302)
		return
	}
	this.Data["Topic"] = topic
	this.Data["tId"] = tId //this.Ctx.Input.Param("0")
	this.Data["admin"] = checkAccount(this.Ctx)
}

func (this *TopicController) Modify() {
	this.TplName = "topic_modify.html"

	tId := this.Input().Get("tid")

	topic, err := models.GetTopic(tId)
	if err != nil {
		beego.Error(err)
		this.Redirect("/", 302)
		return
	}

	this.Data["Topic"] = topic
	this.Data["tId"] = tId
}

func (this *TopicController) Delete() {
	if !checkAccount(this.Ctx) {
		this.Redirect("/login", 302)
		return
	}

	tId := this.Input().Get("tid")
	if "0" != tId {
		models.DelTopic(tId)
	}
	this.Redirect("/topic", 302)
	return
}
