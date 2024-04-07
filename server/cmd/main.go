package main

import (
	"context"
	"fmt"

	logger "github.com/kaungmyathan22/golang-blog/cmd/common"
	"github.com/kaungmyathan22/golang-blog/cmd/database"
)

type Config struct {
	PORT         string
	DATABASE_URL string
	ENV          string
}

type Application struct {
	config Config
	db     *database.Queries
}

func main() {
	app := &Application{
		config: Config{},
	}
	ctx := context.Background()
	conn := app.BootstrapApp()
	defer conn.Close(ctx)
	logger.Info(fmt.Sprintf("server is running at http://localhost:%s", app.config.PORT))
	err := app.Serve()
	if err != nil {
		logger.Error(err.Error())
		logger.Error("server quit")
	}
}
