package models

import (
	"time"
)

type Post struct {
	ID        int64     `json:"id" gorm:"primarykey"`
	Title     string    `json:"title"`
	Content   string    `json:"content"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
	UserID    int64     `json:"userId"`
	User      User      `gorm:"constraint:OnDelete:CASCADE;" json:"-"`
}
