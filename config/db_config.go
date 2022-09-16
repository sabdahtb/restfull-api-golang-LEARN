package config

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
	"github.com/sabdahtb/entity"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// SetupDBConnection is create connection to Mysql DB
func SetupDBConnection() *gorm.DB {
	errENV := godotenv.Load()
	if errENV != nil {
		panic("Failed to load .env File")
	}

	dbUser := os.Getenv("DB_USER")
	dbPass := os.Getenv("DB_PASS")
	dbName := os.Getenv("DB_NAME")
	dbHost := os.Getenv("DB_HOST")

	dsn := fmt.Sprintf("%s:%s@tcp(%s:3306)/%s?charset=utf8mb4&parseTime=True&loc=Local", dbUser, dbPass, dbHost, dbName)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		fmt.Println((err))
		panic("Failed connect to DB")
	}

	db.AutoMigrate(&entity.Book{}, &entity.User{})

	return db
}

// CloseDBConnection is create close connection with DB
func CloseDBConnection(db *gorm.DB) {
	dbSQL, err := db.DB()

	if err != nil {
		panic("Failed to close connection with DB")
	}

	dbSQL.Close()
}
