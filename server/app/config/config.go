package config

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/kaungmyathan22/golang-nextjs-blog/app/database"
	_ "github.com/lib/pq"
	"github.com/spf13/viper"
)

type Config struct {
	PORT   string
	DB_URL string
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
}

func ConnectToDatabase() error {
	const postgresURL = "postgres://admin:admin@localhost:5432/sqlc?sslmode=disable"
	db, err := sql.Open("postgres", postgresURL)
	if err != nil {
		log.Println("Something went wrong while connecting to database!!")
		panic(err)
	}
	defer db.Close()
	err = db.Ping()
	if err != nil {
		log.Println("Something went wrong while pinging to database!!")
		panic(err)
	}
	// ctx := context.Background()
	database.New(db)
	// queries := database.New(db)
	return nil
}
