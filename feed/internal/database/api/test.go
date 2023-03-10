package api

import (
	"feed/internal/service"

	"gorm.io/gorm"
)

func create(db *gorm.DB) error {
	video := service.FeedVideo{
		PlayUrl:  "Test",
		CoverUrl: "Test",
	}
	db.Create(&video)
	return nil
}
func retrieve(db *gorm.DB) error {
	var videosList []*service.FeedVideo
	db.Find(&videosList)
	return nil
}
func update(db *gorm.DB) error {
	testVideo := service.FeedVideo{
		PlayUrl:  "Test",
		CoverUrl: "Test",
	}
	db.Model(&testVideo).Where("play_url = ?", "Test").Update("cover_url", "Test update")
	return nil
}
func delete(db *gorm.DB) error {
	testVideo := service.FeedVideo{
		PlayUrl:  "Test",
		CoverUrl: "Test",
	}
	db.Delete(&testVideo, "play_url = ?", "Test")
	return nil
}

func TestDateBase(db *gorm.DB) {
	if err := create(db); err != nil {
		panic(err)
	}
	if err := retrieve(db); err != nil {
		panic(err)
	}
	if err := update(db); err != nil {
		panic(err)
	}
	if err := delete(db); err != nil {
		panic(err)
	}
}
