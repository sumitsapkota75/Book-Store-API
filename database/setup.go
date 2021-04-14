package database

import (
	"example.com/sumit/models"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

var DB *gorm.DB

func ConnectDatabase() {
	db, err := gorm.Open("sqlite3", "test.db")
	if err != nil {
		panic("Failed to connect to database")
	}
	db.AutoMigrate(&models.Book{})
	DB = db
}
