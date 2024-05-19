package services

import "github.com/kaungmyathan22/golang-nextjs-blog/app/interfaces"

type AuthServiceImpl struct{}

func NewAuthServiceImpl() *AuthServiceImpl {
	return &AuthServiceImpl{}
}

func (svc *AuthServiceImpl) Login(payload interfaces.LoginPayload) (*interfaces.LoginResult, error) {
	return &interfaces.LoginResult{}, nil
}
