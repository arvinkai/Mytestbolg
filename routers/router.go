package routers

import (
	"Myblog/controllers"

	"github.com/astaxie/beego"
)

func init() {

	beego.Router("/", &controllers.MainController{})
	beego.Router("/login", &controllers.LoginController{})
	beego.Router("/exit", &controllers.ExitController{})
	beego.Router("/category", &controllers.CategoryController{})
}
