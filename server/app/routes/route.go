package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/kaungmyathan22/golang-nextjs-blog/app/controllers"
)

func SetupRoute(r *gin.Engine) {
	apiV1Group := r.Group("/api/v1")

	authHandler := controllers.NewAuthControllerImpl()
	authRoute := apiV1Group.Group("/authentication")
	authRoute.POST("/register", authHandler.Register)

	authRoute.POST("/login", authHandler.Login)
	authRoute.POST("/refresh-token", authHandler.RefreshToken)
	authRoute.POST("/forgot-password", authHandler.ForgotPassword)
	authRoute.POST("/reset-password", authHandler.ResetPassword)

	authRoute.POST("/me", authHandler.Me)

	apiV1Group.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

}
