package story

import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"

	"cat-story/internal/repo/models"
)

func Get() ([]models.Story, error) {
	var stories []models.Story

	db, err := gorm.Open(sqlite.Open("storage/main.db"), &gorm.Config{})
	if err != nil{
		return stories, err
	}

	if err := db.Find(&stories).Error; err != nil{
		return stories, err
	}

	return stories, nil
}
