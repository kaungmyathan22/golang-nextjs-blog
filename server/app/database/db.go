package database

import (
	"gorm.io/driver/postgres"

	"github.com/kaungmyathan22/golang-nextjs-blog/app/config"
	"github.com/kaungmyathan22/golang-nextjs-blog/app/logger"
	models "github.com/kaungmyathan22/golang-nextjs-blog/app/models/domain"
	_ "github.com/lib/pq"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectToDatabase() error {
	dsn := config.ConfigInstance.DB_URL
	// Open a connection to the database
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return err
	}
	logger.Info("successfully connected to database.")
	// Migrate the schema
	if err := db.AutoMigrate(&models.User{}); err != nil {
		return err
	}
	if err := db.AutoMigrate(&models.Account{}); err != nil {
		return err
	}
	logger.Info("successfully migrated db schema.")
	DB = db
	return nil
}
