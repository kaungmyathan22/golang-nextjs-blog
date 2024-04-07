package handlers

import (
	"context"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	logger "github.com/kaungmyathan22/golang-blog/cmd/common"
	"github.com/kaungmyathan22/golang-blog/cmd/database"
)

type AuthHandler struct {
	Repository *database.Queries
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
	ctx, cancel := context.WithTimeout(c, time.Second*1)
	defer cancel()
	authors, err := handler.Repository.GetUsers(ctx)
	if err != nil {
		logger.Error(err.Error())
		c.JSON(http.StatusInternalServerError, map[string]string{
			"error": "something went wrong while fetching a list of users.",
		})
		return
	}
	c.JSON(200, map[string]any{
		"data": authors,
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
