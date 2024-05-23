package controllers

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	hash "github.com/kaungmyathan22/golang-nextjs-blog/app"
	"github.com/kaungmyathan22/golang-nextjs-blog/app/database"
	"github.com/kaungmyathan22/golang-nextjs-blog/app/logger"
	"github.com/kaungmyathan22/golang-nextjs-blog/app/models/apis"
	models "github.com/kaungmyathan22/golang-nextjs-blog/app/models/domain"
	jwt "github.com/kaungmyathan22/golang-nextjs-blog/app/utils"
	"gorm.io/gorm"
)

type AuthControllerImpl struct {
}

func NewAuthControllerImpl() *AuthControllerImpl {
	return &AuthControllerImpl{}
}

func (ctrl *AuthControllerImpl) Login(c *gin.Context) {
	var payload *apis.LoginPayload
	if err := c.ShouldBind(&payload); err != nil {
		fmt.Println(err.Error())
		c.JSON(http.StatusBadRequest, apis.APIResponse{
			Status:  http.StatusBadRequest,
			Message: "StatusBadRequest",
			Data: map[string]string{
				"error": "invalid body payload.",
			},
		})
		return
	}
	var user models.User
	result := database.DB.First(&user, "email = ?", payload.Email)
	if result.Error != nil {
		if result.Error == gorm.ErrRecordNotFound {
			c.JSON(http.StatusBadRequest, apis.APIResponse{
				Status:  http.StatusBadRequest,
				Message: "StatusBadRequest",
				Data: map[string]string{
					"message": "invalid email / password.",
				},
			})
		} else {
			c.JSON(http.StatusInternalServerError, apis.APIResponse{
				Status:  http.StatusInternalServerError,
				Message: "StatusInternalServerError",
				Data: map[string]string{
					"message": "Something went wrong",
				},
			})
		}
		return
	}
	if err := hash.ComparePasswordAndHash(payload.Password, user.Password); err != nil {
		c.JSON(http.StatusBadRequest, apis.APIResponse{
			Status:  http.StatusBadRequest,
			Message: "StatusBadRequest",
			Data: map[string]string{
				"message": "invalid email / password.",
			},
		})
		return
	}
	token, err := jwt.SignJwtAuthenticationToken(int(user.ID))
	if err != nil {
		fmt.Println(err)
		c.JSON(http.StatusInternalServerError, apis.APIResponse{
			Status:  http.StatusInternalServerError,
			Message: "StatusInternalServerError",
			Data: map[string]string{
				"message": "Something went wrong",
			},
		})
		return
	}
	c.JSON(http.StatusOK, apis.APIResponse{
		Status:  http.StatusOK,
		Message: "success",
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
			Message: "StatusBadRequest",
			Data: map[string]any{
				"message": "invalid request payload.",
			},
		}
		c.JSON(http.StatusBadRequest, response)
		return
	}
	hashedPassword, err := hash.HashPassword(payload.Password)
	payload.Password = hashedPassword
	if err != nil {
		logger.Error(err.Error())
		c.JSON(http.StatusInternalServerError, apis.APIResponse{
			Status:  http.StatusInternalServerError,
			Message: "StatusInternalServerError",
			Data: map[string]any{
				"message": "something went wrong",
			},
		})
		return
	}
	user := models.User{Name: payload.Name, Password: payload.Password, Email: payload.Email}
	result := database.DB.Create(&user)
	if err := result.Error; err != nil {
		if strings.Contains(err.Error(), "23505") {
			c.JSON(http.StatusConflict, apis.APIResponse{
				Status:  http.StatusConflict,
				Message: "StatusConflict",
				Data: map[string]string{
					"message": "user with given email address already existed.",
				},
			})
		} else {
			logger.Error(err.Error())
			c.JSON(http.StatusInternalServerError, apis.APIResponse{
				Status:  http.StatusInternalServerError,
				Message: "StatusInternalServerError",
				Data: map[string]string{
					"message": "something went wrong",
				},
			})
		}
		return
	}
	c.JSON(http.StatusOK, apis.APIResponse{
		Status:  http.StatusOK,
		Message: "success",
		Data: map[string]any{
			"message": "Registration successful. Please login!",
		},
	})
}

func (ctrl *AuthControllerImpl) Me(c *gin.Context) {
	user, ok := c.Get("user")
	if !ok {
		c.JSON(http.StatusUnauthorized, apis.APIResponse{
			Status:  http.StatusUnauthorized,
			Message: "unauthorized",
			Data:    nil,
		})
		return
	}
	c.JSON(http.StatusOK, apis.APIResponse{
		Status:  http.StatusOK,
		Message: "success",
		Data:    user,
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
