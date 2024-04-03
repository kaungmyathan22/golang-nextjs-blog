package main

import (
	"fmt"
	"log"
	"os"

	logger "github.com/kaungmyathan22/golang-blog/cmd/common"
	"github.com/spf13/viper"
)

var ConfigInstance Config

func (app *Application) LoadConfig() {
	dir, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}
	envFile := fmt.Sprintf("%s/.env", dir)
	viper.SetConfigFile(envFile)
	viper.ReadInConfig()
	viper.AutomaticEnv()
	if err := viper.Unmarshal(&ConfigInstance); err != nil {
		logger.Fatal(err.Error())
	}
	app.config = ConfigInstance
	logger.Info("successfully loaded environment variables")
}
