package main

import (
	"fmt"
	"log"

	"github.com/robfig/cron/v3"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

const (
	host         = "localhost"
	port         = 5432
	databaseName = "mydatabase"
	username     = "myuser"
	password     = "mypassword"
)

func main() {
	// Configure your PostgreSQL database details here
	dsn := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		host, port, username, password, databaseName)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect to database")
	}

	fmt.Println("Connected to database", db)

	db.AutoMigrate(&ExampleModel{})
	fmt.Println("Database migration completed!")

	c := cron.New()
	_, err = c.AddFunc("@every 1m", func() {
		go task(db)
	})

	if err != nil {
		log.Fatal("Error scheduling a task:", err)
	}

	c.Start()

	// Block the main thread as the cron job runs in the background
	select {}
}
