package middlewares

import (
	"fmt"
	"net/http"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/kaungmyathan22/golang-nextjs-blog/app/database"
	"github.com/kaungmyathan22/golang-nextjs-blog/app/logger"
	"github.com/kaungmyathan22/golang-nextjs-blog/app/models/apis"
	models "github.com/kaungmyathan22/golang-nextjs-blog/app/models/domain"
	jwt "github.com/kaungmyathan22/golang-nextjs-blog/app/utils"
)

func IsAuthenticated() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		tokenString := ctx.GetHeader("Authorization")
		logger.Info(fmt.Sprintf("token string is %s", tokenString))
		if tokenString == "" {
			response := apis.APIResponse{
				Status:  http.StatusUnauthorized,
				Message: "Unauthorized",
				Data:    nil,
			}
			ctx.JSON(http.StatusUnauthorized, response)
			ctx.Abort()
			return
		}
		token := strings.Split(tokenString, "Bearer ")[1]
		claims, err := jwt.ValidateJwtAuthenticationToken(token)
		if err != nil {
			logger.Error(err.Error())
			response := apis.APIResponse{
				Status:  http.StatusUnauthorized,
				Message: "Unauthorized",
				Data:    nil,
			}
			ctx.JSON(http.StatusUnauthorized, response)
			ctx.Abort()
			return
		}

		var user *models.User
		result := database.DB.First(&user, claims.Sub)
		if result.Error != nil {
			response := apis.APIResponse{
				Status:  http.StatusUnauthorized,
				Message: "Unauthorized",
				Data:    nil,
			}
			ctx.JSON(http.StatusUnauthorized, response)
			ctx.Abort()
			return
		}
		ctx.Set("user", user)
		ctx.Next()
	}
}

func IsPasswordResetTokenValid() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		bearerToken := ctx.GetHeader("Authorization")
		if bearerToken == "" {
			ctx.JSON(http.StatusUnauthorized, apis.UnauthorizedResponse)
			ctx.Abort()
			return
		}
		tokenString := strings.Split(bearerToken, "Bearer ")[1]
		if tokenString == "" {
			ctx.JSON(http.StatusUnauthorized, apis.GetUnauthorizedResponse("invalid token."))
			ctx.Abort()
			return
		}
		var token *models.Token
		result := database.DB.First(&token, "token = ?", tokenString)
		if err := result.Error; err != nil {
			ctx.JSON(http.StatusUnauthorized, apis.GetUnauthorizedResponse("invalid token."))
			ctx.Abort()
			return
		}
		if token.ExpiredAt.Unix() < time.Now().Unix() {
			ctx.JSON(http.StatusUnauthorized, apis.GetUnauthorizedResponse("token has expired."))
			ctx.Abort()
			return
		}
		var user *models.User
		result = database.DB.First(&user, token.UserID)
		if result.Error != nil {
			logger.Error("user belonging to the given token not found in the database.")
			ctx.JSON(http.StatusUnauthorized, apis.GetUnauthorizedResponse("invalid token."))
			ctx.Abort()
			return
		}
		ctx.Set("user", user)
		ctx.Next()
	}
}
