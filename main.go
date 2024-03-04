package main

import (
	"fmt"
	"golang-api/database"
	"golang-api/models"
	"net/http"
)

func main() {
	db, err := database.InitDB()
	if err != nil {
		panic("failed to connect")
	}

	db.AutoMigrate(&models.User{}, &models.Photo{})

	http.ListenAndServe(":8080", nil)

	fmt.Println("Server on port 8080")
}