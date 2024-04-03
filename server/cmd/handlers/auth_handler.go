package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

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

func (handler *AuthHandler) LoginHandler(c *gin.Context) {
	c.JSON(200, map[string]any{
		"path": "/login",
	})
}

func (handler *AuthHandler) RegisterHandler(c *gin.Context) {
	// validate payload
	// hash user password
	// save user to db.
	c.JSON(200, map[string]any{
		"path": "/register",
	})
}
