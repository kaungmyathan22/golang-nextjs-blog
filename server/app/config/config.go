package config

import (
	"fmt"

	"github.com/kaungmyathan22/golang-nextjs-blog/app/logger"
	_ "github.com/lib/pq"
	"github.com/spf13/viper"
)

type Config struct {
	PORT                   string
	DB_URL                 string
	JWT_TOKEN_SECRET       string
	JWTExpirationInSeconds int
}

var ConfigInstance *Config

func LoadConfig() {
	viper.SetConfigFile("app/config/database.env")
	viper.ReadInConfig()

	// Read the configuration file
	if err := viper.ReadInConfig(); err != nil {
		fmt.Printf("Error reading config file, %s", err)
		return
	}

	// Unmarshal the config into the struct
	if err := viper.Unmarshal(&ConfigInstance); err != nil {
		fmt.Printf("Unable to decode into struct, %v", err)
		return
	}
	logger.Init()
}
