package main

import (
	"fmt"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type User struct {
	Id    uint
	Name  string
	Email string
}

func main() {
	fmt.Println("GORM starting....")

	db, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	err2 := db.AutoMigrate(&User{})
	if err2 != nil {
		return
	}

	db.Create(&User{Name: "Ander", Email: "ander@gmail.com"})

	var user User
	db.First(&user, 1)
	db.First(&user, "name = ?", "Ander")

	db.Model(&user).Update("name", "AnderMe")

	db.Delete(&user, 1)
}
