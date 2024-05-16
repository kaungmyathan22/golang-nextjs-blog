package handlers

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx"
	"github.com/jackc/pgx/v5/pgtype"
	logger "github.com/kaungmyathan22/golang-blog/cmd/common"
	"github.com/kaungmyathan22/golang-blog/cmd/database"
	"github.com/kaungmyathan22/golang-blog/cmd/dto"
	models "github.com/kaungmyathan22/golang-blog/cmd/dto"
	"github.com/kaungmyathan22/golang-blog/cmd/utils"
	"golang.org/x/crypto/bcrypt"
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
	var payload models.RegisterPayload
	if err := c.BindJSON(&payload); err != nil {
		fmt.Println(err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "Failed to parse request body"})
		return
	}
	author, err := handler.Repository.GetUserbyEmail(ctx, payload.Email)
	if err != nil {
		fmt.Println(err.Error())
		logger.Error(err.Error())
		c.JSON(http.StatusInternalServerError, map[string]string{
			"error": "something went wrong while fetching a list of users.",
		})
		return
	}
	c.JSON(200, map[string]any{
		"data": author,
	})
}

func (handler *AuthHandler) RegisterHandler(c *gin.Context) {
	var payload models.RegisterPayload

	if err := c.BindJSON(&payload); err != nil {
		fmt.Println(err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "Failed to parse request body"})
		return
	}

	// hash user password
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(payload.Password), bcrypt.DefaultCost)
	if err != nil {
		fmt.Println("Error hashing password:", err)
		return
	}

	// save user to db.
	dbPayload := database.CreateUserParams{
		Fullname: pgtype.Text{
			String: payload.FullName,
			Valid:  true,
		},
		Email:    payload.Email,
		Password: string(hashedPassword),
	}
	user, err := handler.Repository.CreateUser(c, dbPayload)
	if err != nil {
		fmt.Println(err.Error())
		pgErr := pgx.PgError{}
		fmt.Println(errors.As(err, &pgErr))
		if utils.IsDuplicateKeyError(err) {
			c.JSON(http.StatusConflict, map[string]string{
				"message": "user with given email address already exists.",
			})
			return
		}

		c.JSON(http.StatusInternalServerError, map[string]string{
			"message": err.Error(),
		})
		return
	}
	userDTO := &dto.UserDTO{}
	userDTO.FromUserModel(&user)
	c.JSON(200, userDTO)
}
