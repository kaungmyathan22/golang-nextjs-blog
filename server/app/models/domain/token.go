package models

import (
	"time"
)

type Token struct {
	ID        int64
	Token     string
	Expiredat time.Time
	Tokentype string
	Createdat time.Time
	Updatedat time.Time
}
