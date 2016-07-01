package models

import (
	"time"
)

type Todo struct {
	Id        int
	Name      string
	Completed bool
	Timestamp time.Time
}
