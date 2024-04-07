package dto

import (
	"github.com/jackc/pgx/v5/pgtype"
	"github.com/kaungmyathan22/golang-blog/cmd/database"
)

type UserDTO struct {
	ID       int64       `json:"id"`
	Email    string      `json:"email"`
	Fullname pgtype.Text `json:"fullName"`
	Password string      `json:"-"`
}

func (u *UserDTO) FromUserModel(user *database.User) {
	u.ID = user.ID
	u.Email = user.Email
	u.Fullname = user.Fullname
	u.Password = user.Password
}
