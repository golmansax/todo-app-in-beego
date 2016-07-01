package controllers

import (
	"fmt"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	"github.com/golmansax/todo-app-in-beego/todoapp/models"
	"time"
)

type TodoController struct {
	beego.Controller
}

func (this *TodoController) List() {
	o := orm.NewOrm()
	o.Using("default")

	var todos []*models.Todo
	o.QueryTable("todo").All(&todos)

	this.Data["todos"] = todos
	this.TplName = "todo/list.tpl"
}

func (this *TodoController) New() {
	this.TplName = "todo/new.tpl"
}

func (this *TodoController) Create() {
	o := orm.NewOrm()
	o.Using("default")

	todo := new(models.Todo)
	todo.Name = this.GetString("name")
	todo.Timestamp = time.Now()

	_, err := o.Insert(todo)
	if err != nil {
		fmt.Println(err)
	}

	this.Redirect("/", 301)
}
