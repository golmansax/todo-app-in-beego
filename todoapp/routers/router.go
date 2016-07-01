package routers

import (
	"github.com/astaxie/beego"
	"github.com/golmansax/todo-app-in-beego/todoapp/controllers"
)

func init() {
	beego.Router("/", &controllers.TodoController{}, "get:List")
	beego.Router("/new-todo", &controllers.TodoController{}, "get:New")
	beego.Router("/create-todo", &controllers.TodoController{}, "post:Create")
}
