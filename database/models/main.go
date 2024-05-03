package models

import (
	"log"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	EmailAddress string `gorm:"unique"`
	Name         string
	URLs         []URL
}

type URL struct {
	gorm.Model
	LongUrl  string
	ShortUrl string `gorm:"unique"`
	UserID   uint
}

func GetDatabaseClient() *gorm.DB {
	db, err := gorm.Open(sqlite.Open("../../gotrim.db"))
	if err != nil {
		log.Printf("An error occurred while connecting to database: %v", err)
		return db
	}

	db.AutoMigrate(&User{}, &URL{})

	return db
}
