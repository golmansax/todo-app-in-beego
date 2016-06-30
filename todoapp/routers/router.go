package routers

import (
	"github.com/astaxie/beego"
	"github.com/golmansax/todo-app-in-beego/todoapp/controllers"
)

func init() {
	beego.Router("/", &controllers.MainController{})
}
