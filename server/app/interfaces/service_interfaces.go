package interfaces

import models "github.com/kaungmyathan22/golang-nextjs-blog/app/models/apis"

type AuthService interface {
	Login(payload models.LoginPayload) (models.LoginResult, error)
	Register(payload models.RegisterPayload) (models.MessageResponse, error)
	ChangePassword(payload models.ChangePasswordPayload) (string, error)
	ForgotPassword(payload models.ForgotPasswordPayload) (string, error)
	ResetPassword(payload models.ResetPasswordPayload) (string, error)
	RefreshToken() (string, error)
}
