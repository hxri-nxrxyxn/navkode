package models

import "gorm.io/gorm"

type Landmark struct {
	LandmarkID uint   `gorm:"primaryKey;autoIncrement" json:"id"`
	Name       string `gorm:"varchar(255);not null" json:"name"`
	Location   string `gorm:"varchar(255);not null" json:"location"`
	Star       uint   `gorm:"int;not null" json:"star"`
	Preview    string `gorm:"varchar(255);not null" json:"preview"`
	Category   string `gorm:"varchar(255);not null" json:"category"`
}

func MigrateLandmark(db *gorm.DB) error {
	err := db.AutoMigrate(&Landmark{})
	return err
}
