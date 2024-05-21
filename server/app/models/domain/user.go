package models

import (
	"time"

	"gorm.io/gorm"
)

type User struct {
	ID        int64  `json:"id" gorm:"primarykey"`
	Name      string `json:"name"`
	Password  string `json:"-"`
	Email     string `gorm:"unique" json:"email"`
	Accounts  []Account
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}

type Account struct {
	ID          int64
	Lastloginat time.Time
	UserID      int64
	User        User `gorm:"constraint:OnDelete:CASCADE;"`
	gorm.Model
}
