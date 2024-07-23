package main

import (
	"fmt"
	"log"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type User struct {
	ID    uint `gorm:"primaryKey"`
	Name  string
	Email string
}

func main() {
	fmt.Println("GORM starting...")

	db, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
	if err != nil {
		log.Fatalf("failed to connect database: %v", err)
	}

	if err := db.AutoMigrate(&User{}); err != nil {
		log.Fatalf("failed to migrate database schema: %v", err)
	}

	users := []User{
		{Name: "Ander", Email: "ander@gmail.com"},
		{Name: "Alex", Email: "alex@gmail.com"},
		{Name: "Bob", Email: "bob@gmail.com"},
		{Name: "Charlie", Email: "charlie@gmail.com"},
	}

	for _, user := range users {
		if err := db.Create(&user).Error; err != nil {
			log.Fatalf("failed to create user %s: %v", user.Name, err)
		}
	}

	var user User
	if err := db.First(&user, 2).Error; err != nil {
		log.Fatalf("failed to find user with ID 2: %v", err)
	}
	fmt.Printf("Retrieved user: %+v\n", user)

	if err := db.Where("ID > ?", 10).Delete(&User{}).Error; err != nil {
		log.Fatalf("failed to delete users with ID > 10: %v", err)
	}

	fmt.Println("Deleted all users with ID > 10")

	var userToUpdate User
	if err := db.First(&userToUpdate, 1).Error; err != nil {
		log.Fatalf("failed to find user with ID 1: %v", err)
	}
	userToUpdate.Email = "updated_email@gmail.com"
	if err := db.Save(&userToUpdate).Error; err != nil {
		log.Fatalf("failed to update user: %v", err)
	}
	fmt.Println("Updated user email")

}
