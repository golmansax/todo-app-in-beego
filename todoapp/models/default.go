package models

import (
	"github.com/astaxie/beego/orm"
	"time"
)

type Todo struct {
	Id        int
	Name      string
	Completed bool
	Timestamp time.Time
}

func init() {
	orm.RegisterModel(new(Todo))
}
