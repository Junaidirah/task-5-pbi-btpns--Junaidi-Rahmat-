package database

import (
	"fmt"
	"golang-api/models"
	"log"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var (
	host     = "localhost"
	user     = "mysql"
	password = "root"
	dbPort   = "127.0.0.1:3306"
	dbname   = "databank"
	db       *gorm.DB
	err      error
)

func StartDB() {
	config := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable", host, user, password, dbname, dbPort)
	db, err := gorm.Open(mysql.Open(config), &gorm.Config{})

	if err != nil {
		log.Fatal("Error connecting to database : ", err)
	}

	fmt.Println("Connection success to database")
	db.Debug().AutoMigrate(models.User{}, models.Photo{})
}

func GetDB() *gorm.DB {
	return db
}
