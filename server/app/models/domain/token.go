package models

import (
	"time"
)

// Define a custom type for the enum
type TokenType string

const (
	PasswordReset TokenType = "PasswordReset"
	Refresh       TokenType = "Refresh"
)

type Token struct {
	ID        int64 `gorm:"primaryKey"`
	Token     string
	ExpiredAt time.Time
	Tokentype TokenType `gorm:"type:string;size:50"`
	UserID    int64
	User      User `gorm:"constraint:OnDelete:CASCADE;"`
	Createdat time.Time
	Updatedat time.Time
}
