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
	db.Create(&User{Name: "Alex", Email: "alex@gmail.com"})
	db.Create(&User{Name: "Bob", Email: "bob@gmail.com"})
	db.Create(&User{Name: "Charlie", Email: "charlie@gmail.com"})

	var user User
	db.First(&user, 2)

}
