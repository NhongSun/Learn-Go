package main

import (
	"fmt"

	"gorm.io/gorm"
)

// gorm.Model definition
// type Model struct {
// 	ID        uint           `gorm:"primaryKey"`
// 	CreatedAt time.Time
// 	UpdatedAt time.Time
// 	DeletedAt gorm.DeletedAt `gorm:"index"`
//   }

// in GORM convention, first letter of column must be capital
type Book struct {
	gorm.Model         // from gorm.Model
	Name        string `json:"name"`
	Author      string `json:"author"`
	Description string `json:"description"`
	Price       int    `json:"price"`
}

func getBook(db *gorm.DB, id uint) (*Book, error) {
	var book Book

	result := db.First(&book, id)
	fmt.Println("Book not found")

	if result.Error != nil {
		fmt.Printf("Error querying database: %v\n", result.Error)
		return nil, result.Error
	}
	fmt.Println("Successfully finding book")

	return &book, nil
}

func getBooks(db *gorm.DB) []Book {
	var books []Book

	result := db.Find(&books)

	if result.Error != nil {
		fmt.Printf("Error finding books, %v", result.Error)
	}
	fmt.Println("Successfully finding books")

	return books
}

func createBook(db *gorm.DB, book *Book) error {
	result := db.Create(book)

	if result.Error != nil {
		fmt.Printf("Error creating book, %v", result.Error)
		return result.Error
	}
	fmt.Println("Successfully create new book")

	return nil
}

func updateBook(db *gorm.DB, book *Book) error {
	result := db.Model(&book).Updates(book)

	if result.Error != nil {
		fmt.Printf("Error updating book, %v", result.Error)
		return result.Error
	}
	fmt.Println("Successfully update book")

	return nil
}

// This is a SOFT DELETE
// If model includes a gorm.DeletedAt field
// , it will get soft delete ability automatically
// the record WON’T be removed from the database
// , but GORM will set the DeletedAt‘s value to the current time
// and the data is not findable with GORM's Query methods anymore
func deleteBook(db *gorm.DB, id uint) error {
	var book Book

	result := db.Delete(&book, id)

	if result.Error != nil {
		fmt.Printf("Error deleting book, %v", result.Error)
		return result.Error
	}
	fmt.Println("Successfully delete book")

	return nil
}

// func searchBook(db *gorm.DB, boookName string) *Book {
// 	var book Book

// 	// search by name, order by price and limit 1
// 	result := db.Where("name = ?", boookName).Order("price").First(&book)

// 	if result.Error != nil {
// 		fmt.Printf("Error searching book, %v", result.Error)
// 	}
// 	fmt.Println("Successfully search book")

// 	return &book
// }

// func searchBooks(db *gorm.DB, boookName string) []Book {
// 	var books []Book

// 	result := db.Where("name = ?", boookName).Order("price").Find(&books)

// 	if result.Error != nil {
// 		fmt.Printf("Error searching book, %v", result.Error)
// 	}
// 	fmt.Println("Successfully search book")

// 	// slice normally give pointer
// 	return books
// }
