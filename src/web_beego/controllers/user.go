package controllers

import "github.com/astaxie/beego"

type UserController struct {
	beego.Controller
}

func (c *UserController) Get() {
	c.Data["Website"] = "beego.me0000000000"
	c.Data["Email"] = "astaxie@gmail.com"
	c.TplName = "index.tpl"
}
