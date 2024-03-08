package main

import (
	"golang-api/database"
	"golang-api/router"
)

func main() {
	database.StartDB()
	r := router.StartApp()
	r.Run(":8080")
}
