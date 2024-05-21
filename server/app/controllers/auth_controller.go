package controllers

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	hash "github.com/kaungmyathan22/golang-nextjs-blog/app"
	"github.com/kaungmyathan22/golang-nextjs-blog/app/database"
	"github.com/kaungmyathan22/golang-nextjs-blog/app/logger"
	"github.com/kaungmyathan22/golang-nextjs-blog/app/models/apis"
	models "github.com/kaungmyathan22/golang-nextjs-blog/app/models/domain"
	"github.com/kaungmyathan22/golang-nextjs-blog/app/services"
	jwt "github.com/kaungmyathan22/golang-nextjs-blog/app/utils"
	"gorm.io/gorm"
)

type AuthControllerImpl struct {
	SVC *services.AuthServiceImpl
}

func NewAuthControllerImpl(svc *services.AuthServiceImpl) *AuthControllerImpl {
	return &AuthControllerImpl{SVC: svc}
}

func (ctrl *AuthControllerImpl) Login(c *gin.Context) {
	var payload *apis.LoginPayload
	if err := c.ShouldBind(&payload); err != nil {
		fmt.Println(err.Error())
		c.JSON(http.StatusBadRequest, apis.APIResponse{
			Status:  http.StatusBadRequest,
			Message: "invalid body payload.",
		})
		return
	}
	var user models.User
	result := database.DB.First(&user, "email = ?", payload.Email)
	if result.Error != nil {
		if result.Error == gorm.ErrRecordNotFound {
			c.JSON(http.StatusBadRequest, apis.APIResponse{
				Status:  http.StatusBadRequest,
				Message: "invalid email / password.",
			})
		} else {
			c.JSON(http.StatusInternalServerError, apis.APIResponse{
				Status:  http.StatusInternalServerError,
				Message: "something went wrong",
			})
		}
		return
	}
	if err := hash.ComparePasswordAndHash(payload.Password, user.Password); err != nil {
		c.JSON(http.StatusBadRequest, apis.APIResponse{
			Status:  http.StatusBadRequest,
			Message: "invalid email / password.",
		})
		return
	}
	token, err := jwt.SignJwtAuthenticationToken(int(user.ID))
	if err != nil {
		fmt.Println(err)
		c.JSON(http.StatusBadRequest, apis.APIResponse{
			Status:  http.StatusBadRequest,
			Message: "error signing token.",
		})
		return
	}
	c.JSON(http.StatusOK, apis.APIResponse{
		Status: http.StatusOK,
		Data: map[string]any{
			"token": token,
			"user":  user,
		},
	})
}

func (ctrl *AuthControllerImpl) Register(c *gin.Context) {
	// ctrl.SVC.Register()
	var payload *apis.RegisterPayload
	if err := c.ShouldBind(&payload); err != nil {
		fmt.Println(err)
		response := apis.APIResponse{
			Status:  http.StatusBadRequest,
			Message: "invalid request payload.",
		}
		c.JSON(http.StatusBadRequest, response)
		return
	}
	hashedPassword, err := hash.HashPassword(payload.Password)
	payload.Password = hashedPassword
	if err != nil {
		c.JSON(http.StatusInternalServerError, apis.APIResponse{
			Status:  http.StatusInternalServerError,
			Message: "error while hashing password",
		})
		return
	}
	if err = ctrl.SVC.Register(payload); err != nil {
		logger.Error(err.Error())
		c.JSON(http.StatusConflict, apis.APIResponse{
			Status:  http.StatusConflict,
			Message: err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, apis.APIResponse{
		Status:  http.StatusOK,
		Message: "Registration successful. Please login!",
	})
}

func (ctrl *AuthControllerImpl) ChangePassword(c *gin.Context) {
	c.JSON(200, map[string]string{
		"message": "ChangePassword",
	})
}

func (ctrl *AuthControllerImpl) ForgotPassword(c *gin.Context) {
	c.JSON(200, map[string]string{
		"message": "ForgotPassword",
	})
}

func (ctrl *AuthControllerImpl) ResetPassword(c *gin.Context) {
	c.JSON(200, map[string]string{
		"message": "ResetPassword",
	})
}

func (ctrl *AuthControllerImpl) RefreshToken(c *gin.Context) {
	c.JSON(200, map[string]string{
		"message": "RefreshToken",
	})
}
