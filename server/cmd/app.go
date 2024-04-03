package main

import (
	"github.com/gin-gonic/gin"
	logger "github.com/kaungmyathan22/golang-blog/cmd/common"
	"github.com/kaungmyathan22/golang-blog/cmd/routes"
)

func (app *Application) BootstrapApp() {
	logger.Init()
	app.LoadConfig()
}

func (app *Application) Serve() error {
	r := gin.Default()
	routes.InitRoutes(r)
	err := r.Run(app.config.PORT)
	if err != nil {
		logger.Error(err.Error())
		logger.Error("server quit")
	}
	return nil
}
