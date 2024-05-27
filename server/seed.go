package main

import (
	"fmt"

	"github.com/kaungmyathan22/golang-nextjs-blog/app/logger"
	models "github.com/kaungmyathan22/golang-nextjs-blog/app/models/domain"
	"gorm.io/gorm"
)

func SeedPostDatabase(db *gorm.DB) {
	for i := 1; i <= 100; i++ {
		db.Create(&models.Post{
			Title:   fmt.Sprintf("Title %d", i),
			Content: fmt.Sprintf("Content %d.com", i),
			UserID:  1,
		})
	}
	logger.Info("successfully seeded posts.")
}
