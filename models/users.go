package models

import "gorm.io/gorm"

type Users struct {
	ID            uint   `gorm:"primaryKey"`
	Username      string `gorm:"unique;not null"`
	Password_Hash string `gorm:"not null"`
	IsAdmin       bool   `gorm:"default:false"`
}

func Migrate(db *gorm.DB) {
	db.AutoMigrate(&Users{})
}
