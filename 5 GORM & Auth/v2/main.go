package main

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

const (
	host         = "localhost"
	port         = 5432
	databaseName = "mydatabase"
	username     = "myuser"
	password     = "mypassword"
)

func main() {
	// Load env file
	if err := godotenv.Load(); err != nil {
		log.Fatal("Error loading env file")
	}

	// Configure your PostgreSQL database details here
	dsn := fmt.Sprintf("host=%s port=%d user=%s "+"password=%s dbname=%s sslmode=disable",
		host, port, username, password, databaseName)

	// New logger for detailed SQL logging
	newLogger := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags), // io writer
		logger.Config{
			SlowThreshold: time.Second, // Slow SQL threshold
			LogLevel:      logger.Info, // Log level: Silent Error Warn Info (show all)
			Colorful:      true,        // Enable color
		},
	)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		Logger: newLogger,
	})

	if err != nil {
		// panic -> break process
		panic("failed to connect to database")
	}

	print(db)
	fmt.Println("Successful connection")

	// ขาสร้าง
	publisher := Publisher{
		Details: "Publisher Details",
		Name:    "Publisher Name",
	}
	_ = createPublisher(db, &publisher)

	// Example data for a new author
	author := Author{
		Name: "Author Name",
	}
	_ = createAuthor(db, &author)

	// // Example data for a new book with an author
	book := Book{
		Name:        "Book Title",
		Author:      "Book Author",
		Description: "Book Description",
		PublisherID: publisher.ID,     // Use the ID of the publisher created above
		Authors:     []Author{author}, // Add the created author
	}
	_ = createBookWithAuthor(db, &book, []uint{author.ID})

	// ขาเรียก

	// Example: Get a book with its publisher
	bookWithPublisher, err := getBookWithPublisher(db, 1) // assuming a book with ID 1
	if err != nil {
		// Handle error
	}

	// Example: Get a book with its authors
	bookWithAuthors, err := getBookWithAuthors(db, 1) // assuming a book with ID 1
	if err != nil {
		// Handle error
	}

	// Example: List books of a specific author
	authorBooks, err := listBooksOfAuthor(db, 1) // assuming an author with ID 1
	if err != nil {
		// Handle error
	}

	fmt.Println(bookWithPublisher)
	fmt.Println(bookWithAuthors)
	fmt.Println(authorBooks)
}
