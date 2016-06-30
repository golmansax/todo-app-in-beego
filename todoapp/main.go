package main

import (
	"github.com/astaxie/beego"
	_ "github.com/golmansax/todo-app-in-beego/todoapp/routers"
)

func main() {
	beego.Run()
}
