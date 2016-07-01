package main

import (
	"fmt"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	"github.com/golmansax/todo-app-in-beego/todoapp/models"
	_ "github.com/golmansax/todo-app-in-beego/todoapp/routers"
	_ "github.com/mattn/go-sqlite3"
)

func init() {
	orm.RegisterDriver("sqlite", orm.DRSqlite)
	orm.RegisterDataBase("default", "sqlite3", "db/todoapp_beego.sqlite3")
	orm.RegisterModel(new(models.Todo))
}

func main() {
	err := orm.RunSyncdb("default", true, true)
	if err != nil {
		fmt.Println(err)
	}

	beego.Run()
}
