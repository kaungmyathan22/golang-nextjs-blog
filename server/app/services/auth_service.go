package services

import (
	"github.com/kaungmyathan22/golang-nextjs-blog/app/models/apis"
)

type AuthServiceImpl struct{}

func NewAuthServiceImpl() *AuthServiceImpl {
	return &AuthServiceImpl{}
}

func (svc *AuthServiceImpl) Login(payload apis.LoginPayload) (*apis.LoginResult, error) {
	return &apis.LoginResult{}, nil
}
