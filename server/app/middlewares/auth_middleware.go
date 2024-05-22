package middlewares

import (
	"fmt"
	"net/http"
	"strings"

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
		result := database.DB.First(&user, claims.ID)
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
