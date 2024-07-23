package main

import (
	"fmt"
	"log"

	"gorm.io/driver/mysql"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type User struct {
	ID    uint `gorm:"primaryKey"`
	Name  string
	Email string
	Age   int
}

func main() {
	fmt.Println("GORM starting...")

	// SQLite configuration
	sqliteDB, err := gorm.Open(sqlite.Open("sqlite_gormdb.db"), &gorm.Config{})
	if err != nil {
		log.Fatalf("failed to connect to SQLite database: %v", err)
	}

	// MySQL configuration
	dsn := "root:123456@tcp(127.0.0.1:3310)/gormdb?charset=utf8mb4&parseTime=True&loc=Local"
	mysqlDB, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("failed to connect to MySQL database: %v", err)
	}

	// Migrate the schema for both databases
	if err := sqliteDB.AutoMigrate(&User{}); err != nil {
		log.Fatalf("failed to migrate SQLite database schema: %v", err)
	}
	if err := mysqlDB.AutoMigrate(&User{}); err != nil {
		log.Fatalf("failed to migrate MySQL database schema: %v", err)
	}

	// Insert data into SQLite
	users := []User{
		{Name: "Ander", Email: "ander@gmail.com"},
		{Name: "Alex", Email: "alex@gmail.com"},
		{Name: "Bob", Email: "bob@gmail.com"},
		{Name: "Charlie", Email: "charlie@gmail.com"},
	}

	for _, user := range users {
		if err := sqliteDB.Create(&user).Error; err != nil {
			log.Fatalf("failed to create user in SQLite %s: %v", user.Name, err)
		}
	}

	// Insert data into MySQL
	for _, user := range users {
		if err := mysqlDB.Create(&user).Error; err != nil {
			log.Fatalf("failed to create user in MySQL %s: %v", user.Name, err)
		}
	}

	// Perform operations on SQLite
	var user User
	if err := sqliteDB.First(&user, 2).Error; err != nil {
		log.Fatalf("failed to find user with ID 2 in SQLite: %v", err)
	}
	fmt.Printf("Retrieved user from SQLite: %+v\n", user)

	// Perform operations on MySQL
	if err := mysqlDB.First(&user, 2).Error; err != nil {
		log.Fatalf("failed to find user with ID 2 in MySQL: %v", err)
	}
	fmt.Printf("Retrieved user from MySQL: %+v\n", user)

	// Other operations can be similarly performed on both databases as needed

	// Transaction example with MySQL
	err = mysqlDB.Transaction(func(tx *gorm.DB) error {
		user1 := User{Name: "User1", Email: "user1@gmail.com"}
		if err := tx.Create(&user1).Error; err != nil {
			return err
		}

		user2 := User{Name: "User2", Email: "user2@gmail.com"}
		if err := tx.Create(&user2).Error; err != nil {
			return err
		}

		return nil
	})
	if err != nil {
		log.Fatalf("MySQL transaction failed: %v", err)
	}
	fmt.Println("MySQL transaction successful")
}
