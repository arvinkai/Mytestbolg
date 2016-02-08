package main

import (
	"Myblog/models"
	_ "Myblog/routers"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
)

func init() {
	models.RegistDB()
}
func main() {
	orm.Debug = true
	orm.RunSyncdb("default", false, true)
	beego.Run()
}
