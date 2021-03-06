package controllers

import (
	"Myblog/models"

	"github.com/astaxie/beego"
)

type MainController struct {
	beego.Controller
}

func (this *MainController) Get() {
	this.Data["IsHome"] = true
	this.Data["IsLogin"] = checkAccount(this.Ctx)
	this.TplName = "home.html"
	topics, err := models.GetAllTopics(this.Input().Get("cate"), true)
	if err != nil {
		beego.Error(err.Error())
	} else {
		this.Data["Topics"] = topics
	}
	category, err := models.GetAllCategories()
	if err != nil {
		beego.Error(err.Error())
	} else {
		this.Data["Categorys"] = category
	}
	return
}
