package models

import "gorm.io/gorm"

type TimeTableItem struct {
	ID        uint   `gorm:"primaryKey;autoIncrement" json:"id"`
	UserID    uint   `json:"userId"`
	TimeStart string `json:"timeStart"`
	TimeEnd   string `json:"timeEnd"`
	Name      string `json:"name"`
	ImgColor  string `json:"imgColor"`
	ImgURL    string `json:"imgUrl"`
	LikeCount int    `json:"likeCount"`
}

func MigrateTimeTableItem(db *gorm.DB) error {
	db.AutoMigrate(&TimeTableItem{})

	return nil
}
