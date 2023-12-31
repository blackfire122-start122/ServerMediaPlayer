package pkg

import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var DB *gorm.DB

type User struct {
	gorm.Model
	Id       uint64 `gorm:"primaryKey"`
	Username string
	Password string
	Email    string
	Phone    string
	Image    string
}

type Song struct {
	gorm.Model
	Id    uint64 `gorm:"primaryKey"`
	Name  string
	File  string
	Image string
}

func init() {
	db, err := gorm.Open(sqlite.Open("database.db"), &gorm.Config{})
	DB = db

	if err != nil {
		panic("failed to connect database")
	}

	err = DB.AutoMigrate(&User{}, &Song{})
	if err != nil {
		panic("Error autoMigrate: ")
	}
}
