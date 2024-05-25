package handlers

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/asaskevich/govalidator"
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
		c.JSON(http.StatusBadRequest, apis.GetStatusBadRequestResponse("invalid body payload."))
		return
	}
	_, err := govalidator.ValidateStruct(payload)
	if err != nil {
		c.JSON(http.StatusBadRequest, apis.GetStatusBadRequestResponse(err.Error()))
		return
	}
	var user models.User
	result := database.DB.First(&user, "email = ?", payload.Email)
	if result.Error != nil {
		if result.Error == gorm.ErrRecordNotFound {
			c.JSON(http.StatusBadRequest, apis.GetStatusBadRequestResponse("invalid email / password."))
		} else {
			c.JSON(http.StatusInternalServerError, apis.InternalServerErrorResponse)
		}
		return
	}
	if err := hash.ComparePasswordAndHash(payload.Password, user.Password); err != nil {
		c.JSON(http.StatusBadRequest, apis.GetStatusBadRequestResponse("invalid email / password."))
		return
	}
	token, err := jwt.SignJwtAuthenticationToken(int(user.ID))
	if err != nil {
		fmt.Println(err)
		c.JSON(http.StatusInternalServerError, apis.InternalServerErrorResponse)
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
	var payload apis.RegisterPayload
	if err := c.ShouldBind(&payload); err != nil {
		logger.Error(err.Error())
		c.JSON(http.StatusBadRequest, apis.GetStatusBadRequestResponse("invalid request payload"))

		return
	}
	_, err := govalidator.ValidateStruct(payload)
	if err != nil {
		c.JSON(http.StatusBadRequest, apis.GetStatusBadRequestResponse(err.Error()))
		return
	}
	hashedPassword, err := hash.HashPassword(payload.Password)
	payload.Password = hashedPassword
	if err != nil {
		logger.Error(err.Error())
		c.JSON(http.StatusInternalServerError, apis.InternalServerErrorResponse)
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
			c.JSON(http.StatusInternalServerError, apis.InternalServerErrorResponse)
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
		c.JSON(http.StatusUnauthorized, apis.UnauthorizedResponse)
		return
	}
	c.JSON(http.StatusOK, apis.APIResponse{
		Status:  http.StatusOK,
		Message: "success",
		Data:    user,
	})
}

func (ctrl *AuthControllerImpl) ChangePassword(c *gin.Context) {
	var payload *apis.ChangePasswordPayload
	if err := c.ShouldBind(&payload); err != nil {
		c.JSON(http.StatusBadRequest, apis.GetStatusBadRequestResponse("invalid request payload."))
		return
	}
	_, err := govalidator.ValidateStruct(payload)
	if err != nil {
		c.JSON(http.StatusBadRequest, apis.GetStatusBadRequestResponse(err.Error()))
		return
	}
	rawUser, ok := c.Get("user")
	if !ok {
		c.JSON(http.StatusUnauthorized, apis.UnauthorizedResponse)
		return
	}
	user, ok := rawUser.(*models.User)
	if !ok {
		logger.Error("error while user type conversion")
		c.JSON(http.StatusInternalServerError, apis.InternalServerErrorResponse)
	}
	if err := hash.ComparePasswordAndHash(payload.OldPassword, user.Password); err != nil {
		c.JSON(http.StatusBadRequest, apis.GetStatusBadRequestResponse("invalid old password"))
		return
	}
	if err := hash.ComparePasswordAndHash(payload.NewPassword, user.Password); err == nil {
		c.JSON(http.StatusBadRequest, apis.GetStatusBadRequestResponse("new password can't be the same with old password"))
		return
	}
	hashedNewPassword, err := hash.HashPassword(payload.NewPassword)
	if err != nil {
		logger.Error("error while hashing new password")
		c.JSON(http.StatusInternalServerError, apis.InternalServerErrorResponse)
		return
	}
	user.Password = hashedNewPassword
	result := database.DB.Save(&user)
	if err := result.Error; err != nil {
		logger.Error("error while saving user new password")
		logger.Error(err.Error())
		c.JSON(http.StatusInternalServerError, apis.InternalServerErrorResponse)
		return
	}
	c.JSON(200, apis.GetStatusAcceptedResponse(map[string]string{
		"message": "successfully updated the password.",
	}))
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
