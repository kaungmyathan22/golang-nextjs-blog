package handlers

import "net/http"

type AuthHandler struct {
}

type IAuthHandler interface {
	LoginHandler() http.Response
	RegisterHandler() http.Response
	EmailConfirmationHandler() http.Response
	ForgotPasswordHandler() http.Response
	ResetPasswordHandler() http.Response
	ChangePasswordHandler() http.Response
	EditProfileHandler() http.Response
	RefreshTokenHandler() http.Response
}
