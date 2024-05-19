package main

import (
	"github.com/gin-gonic/gin"
	"github.com/kaungmyathan22/golang-nextjs-blog/app/config"
	"github.com/kaungmyathan22/golang-nextjs-blog/app/routes"
)

func main() {
	// load env
	config.LoadConfig()
	// setup route
	r := gin.Default()
	r.Use(gin.Logger())
	routes.SetupRoute(r)
	r.Run(config.ConfigInstance.PORT)

}
