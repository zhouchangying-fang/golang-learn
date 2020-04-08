package routers

import (
	"github.com/astaxie/beego"
	"web_beego/controllers"
)

func init() {
	beego.Router("/", &controllers.MainController{})
	beego.Router("/user", &controllers.UserController{})
}
