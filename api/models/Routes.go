package models

import "gorm.io/gorm"

type Routes struct {
	RoutesID uint   `gorm:"primaryKey;autoIncrement" json:"id"`
	Link     string `gorm:"varchar(255);" json:"link"`
	Traffic  int    `gorm:"int" json:"traffic"`
	Safety   int    `gorm:"int" json:"safety"`
	Weather  int    `gorm:"int" json:"weather"`
}

func MigrateRoutes(db *gorm.DB) error {
	err := db.AutoMigrate(&Routes{})
	return err
}
