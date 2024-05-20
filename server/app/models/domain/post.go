package models

import (
	"time"
)

type Category struct {
	ID        int64
	Title     string
	Createdat time.Time
	Updatedat time.Time
}

type Post struct {
	ID         int64
	Title      string
	Content    string
	Createdat  time.Time
	Updatedat  time.Time
	Authorid   int
	Categoryid int
}
