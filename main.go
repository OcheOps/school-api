package main

import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"school-api/models"
)

func main() {
	db, err := gorm.Open(sqlite.Open("school.db"), &gorm.Config{})
	if err != nil {
		panic("Failed to connect to the database")
	}
	
	db.AutoMigrate(&models.Teacher{})
	db.AutoMigrate(&models.Classroom{})
	db.AutoMigrate(&models.Student{})
	
	// Your API routes and server setup go here

	
}
