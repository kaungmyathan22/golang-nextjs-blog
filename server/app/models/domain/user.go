package models

import (
	"time"
)

type User struct {
	ID        int64     `json:"id" gorm:"primarykey"`
	Name      string    `json:"name"`
	Password  string    `json:"-"`
	Email     string    `gorm:"unique" json:"email"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}
