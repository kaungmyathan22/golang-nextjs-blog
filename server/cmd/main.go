package main

import (
	"fmt"

	logger "github.com/kaungmyathan22/golang-blog/cmd/common"
)

type Config struct {
	PORT string
	ENV  string
}

type Application struct {
	config Config
}

func main() {
	app := &Application{
		config: Config{},
	}
	app.BootstrapApp()
	logger.Info(fmt.Sprintf("server is running at http://localhost:%s", app.config.PORT))
	err := app.Serve()
	if err != nil {
		logger.Error(err.Error())
		logger.Error("server quit")
	}
}
