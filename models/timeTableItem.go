package models

import "gorm.io/gorm"

type TimeTableItem struct {
	ID        uint   `gorm:"primaryKey;autoIncrement" json:"id"`
	UserID    uint   `json:"user_id"`
	TimeStart string `json:"time_start"`
	TimeEnd   string `json:"time_end"`
	Name      string `json:"name"`
	ImgColor  string `json:"img_color"`
	ImgURL    string `json:"img_url"`
	LikeCount int    `json:"like_count"`
}

func MigrateTimeTableItem(db *gorm.DB) error {
	db.AutoMigrate(&TimeTableItem{})

	return nil
}
