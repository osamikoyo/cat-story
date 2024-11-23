package story

import (
	"context"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"

	"cat-story/internal/repo/models"
)

func Add(ctx context.Context, story models.Story) error {
	db, err := gorm.Open(sqlite.Open("storage/main.db"), &gorm.Config{})
	if err != nil{
		return err
	}	

	if err := db.Create(&story).Error; err != nil{
		return err
	}

	return nil
}
