package models

import (
	"gorm.io/gorm"
)

type User struct {
	UserID    uint    `gorm:"primaryKey;autoIncrement" json:"id"`
	Name      *string `json:"name"`
	Email     string  `gorm:"unique;not null" json:"email"`
	Password  string  `gorm:"not null" json:"password"`
	Longitude string  `gorm:"varchar(255);" json:"lon"`
	Latitude  string  `gorm:"varchar(255);" json:"lat"`
}

func MigrateUser(db *gorm.DB) error {
	err := db.AutoMigrate(&User{})
	return err
}
