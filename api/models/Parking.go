package models

import "gorm.io/gorm"

type Parking struct {
	ParkingID uint   `gorm:"primaryKey;autoIncrement" json:"id"`
	Name      string `gorm:"varchar(255);" json:"name"`
	Longitude string `gorm:"varchar(255);" json:"lon"`
	Latitude  string `gorm:"varchar(255);" json:"lat"`
}

func MigrateParking(db *gorm.DB) error {
	err := db.AutoMigrate(&Parking{})
	return err
}
