package services

import (
	"fmt"

	"github.com/kaungmyathan22/golang-nextjs-blog/app/database"
	"github.com/kaungmyathan22/golang-nextjs-blog/app/models/apis"
	models "github.com/kaungmyathan22/golang-nextjs-blog/app/models/domain"
)

type AuthServiceImpl struct{}

func NewAuthServiceImpl() *AuthServiceImpl {
	return &AuthServiceImpl{}
}

func (svc *AuthServiceImpl) Register(payload *apis.RegisterPayload) error {
	user := models.User{Name: payload.Name, Password: payload.Password, Email: payload.Email}
	result := database.DB.Create(&user)
	if err := result.Error; err != nil {
		return fmt.Errorf("user with email %s already exists", user.Email)
	}
	return nil
}
