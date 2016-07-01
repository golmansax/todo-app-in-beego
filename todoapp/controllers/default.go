package controllers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	"github.com/golmansax/todo-app-in-beego/todoapp/models"
)

type MainController struct {
	beego.Controller
}

func (c *MainController) Get() {
	o := orm.NewOrm()
	o.Using("default")

	var todos []*models.Todo
	o.QueryTable("todo").All(&todos)

	c.Data["todos"] = todos
	c.TplName = "index.tpl"
}
