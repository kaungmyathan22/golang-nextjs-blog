package main

import (
	"github.com/gin-gonic/gin"
	"github.com/kaungmyathan22/golang-nextjs-blog/app/config"
	"github.com/kaungmyathan22/golang-nextjs-blog/app/database"
	"github.com/kaungmyathan22/golang-nextjs-blog/app/middlewares"
	"github.com/kaungmyathan22/golang-nextjs-blog/app/routes"
)

func main() {
	// load env
	config.LoadConfig()
	err := database.ConnectToDatabase()
	if err != nil {
		panic(err)
	}
	// setup route
	r := gin.Default()
	r.Use(middlewares.LoggerMiddleware())
	routes.SetupRoute(r)
	r.Run(config.ConfigInstance.PORT)

}
