package models

import "gorm.io/gorm"

type Restaurant struct {
	RestaurantID uint   `gorm:"primaryKey;autoIncrement" json:"id"`
	Name      string `gorm:"varchar(255);" json:"name"`
	Longitude string `gorm:"varchar(255);" json:"lon"`
	Latitude  string `gorm:"varchar(255);" json:"lat"`
}

func MigrateRestaurant(db *gorm.DB) error {
	err := db.AutoMigrate(&Restaurant{})
	return err
}
