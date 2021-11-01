package main

import (
	"rest-api-crud-gin/models"
	"rest-api-crud-gin/routes"
)

func main() {

	db := models.SetupDB()
	db.AutoMigrate(&models.Task{})

	r := routes.SetupRoutes(db)
	r.Run()
}
