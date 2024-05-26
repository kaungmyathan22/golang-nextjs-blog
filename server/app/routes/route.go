package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/kaungmyathan22/golang-nextjs-blog/app/handlers"
	"github.com/kaungmyathan22/golang-nextjs-blog/app/middlewares"
)

func SetupRoute(r *gin.Engine) {
	SetupAuthRoutes(r)
	SetupPostRoutes(r)
}

func SetupPostRoutes(r *gin.Engine) {
	apiV1Group := r.Group("/api/v1")
	postRoute := apiV1Group.Group("/posts")
	postHandler := handlers.NewPostsHandlerImpl()
	postRoute.POST("/", middlewares.IsAuthenticated(), postHandler.CreatePost)
	postRoute.POST("/:id", middlewares.IsAuthenticated(), postHandler.UpdatePost)
	postRoute.DELETE("/:id", middlewares.IsAuthenticated(), postHandler.DeletePost)

	postRoute.GET("/:id", postHandler.GetPost)
	postRoute.GET("/", postHandler.GetPosts)
}

func SetupAuthRoutes(r *gin.Engine) {
	apiV1Group := r.Group("/api/v1")
	emailHandler := handlers.NewEmailHandler()
	authHandler := handlers.NewAuthControllerImpl(emailHandler)
	authRoute := apiV1Group.Group("/authentication")
	authRoute.POST("/register", authHandler.Register)

	authRoute.POST("/login", authHandler.Login)
	authRoute.POST("/refresh-token", authHandler.RefreshToken)
	authRoute.POST("/forgot-password", authHandler.ForgotPassword)
	authRoute.POST("/reset-password", middlewares.IsPasswordResetTokenValid(), authHandler.ResetPassword)
	authRoute.GET("/logout", authHandler.Logout)

	authRoute.POST("/change-password", middlewares.IsAuthenticated(), authHandler.ChangePassword)
	authRoute.GET("/me", middlewares.IsAuthenticated(), authHandler.Me)

	apiV1Group.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
}
