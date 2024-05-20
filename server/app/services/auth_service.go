package services

import models "github.com/kaungmyathan22/golang-nextjs-blog/app/models/apis"

type AuthServiceImpl struct{}

func NewAuthServiceImpl() *AuthServiceImpl {
	return &AuthServiceImpl{}
}

func (svc *AuthServiceImpl) Login(payload models.LoginPayload) (*models.LoginResult, error) {
	return &models.LoginResult{}, nil
}
