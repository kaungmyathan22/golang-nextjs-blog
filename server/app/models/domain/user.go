package domain

import "time"

type User struct {
	ID        int64
	Name      string
	Password  string
	Email     string
	Createdat time.Time
	Updatedat time.Time
}

type Account struct {
	ID          int64
	Lastloginat time.Time
	Createdat   time.Time
	Updatedat   time.Time
}
