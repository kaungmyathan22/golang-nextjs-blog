package middlewares

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
	"github.com/kaungmyathan22/golang-nextjs-blog/app/config"
	"github.com/kaungmyathan22/golang-nextjs-blog/app/models/apis"
)

func IsAuthenticated() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		tokenString := ctx.GetHeader("Authorization")

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

		token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
			return []byte(config.ConfigInstance.JWT_TOKEN_SECRET), nil
		})

		if err != nil || !token.Valid {
			response := apis.APIResponse{
				Status:  http.StatusUnauthorized,
				Message: "Unauthorized",
				Data:    nil,
			}
			ctx.JSON(http.StatusUnauthorized, response)
			ctx.Abort()
			return
		}

		ctx.Next()
	}

}
