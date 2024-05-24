package interfaces

import "github.com/kaungmyathan22/golang-nextjs-blog/app/models/apis"

type AuthService interface {
	Login(payload apis.LoginPayload) (apis.LoginResponse, error)
	Register(payload apis.RegisterPayload) (apis.MessageResponse, error)
	ChangePassword(payload apis.ChangePasswordPayload) (string, error)
	ForgotPassword(payload apis.ForgotPasswordPayload) (string, error)
	ResetPassword(payload apis.ResetPasswordPayload) (string, error)
	RefreshToken() (string, error)
}
