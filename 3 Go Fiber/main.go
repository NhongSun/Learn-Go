package main

import (
	"log"
	"os"

	"github.com/gofiber/fiber/v2"
	jwtware "github.com/gofiber/jwt/v2"
	"github.com/gofiber/template/html/v2"
	"github.com/joho/godotenv"
)

type Book struct {
	ID     int    `json:"id"`    // `json:"id"` is a tag for fiber to know that this field is id
	Title  string `json:"title"` // serialize the field to json with the name title
	Author string `json:"author"`
}

var books []Book

func main() {
	// Load env file
	// can be used by os.Getenv("KEY")
	if err := godotenv.Load(); err != nil {
		log.Fatal("Error loading env file")
	}

	// Initialize standard Go html template engine
	// path to the views folder
	engine := html.New("./views", ".html")

	// Like ExpressJs
	// Pass the engine to Fiber
	app := fiber.New(fiber.Config{
		Views: engine,
	})

	// Initializa data
	books = append(books, Book{ID: 1, Title: "1984", Author: "George Orwell"})
	books = append(books, Book{ID: 2, Title: "The Great Gatsby", Author: "F. Scott Fitzgerald"})
	books = append(books, Book{ID: 3, Title: "Computer Architecture", Author: "Krerk Piromposa"})
	books = append(books, Book{ID: 4, Title: "Digital Founding", Author: "True Digital Academy"})

	app.Post("/login", login)

	// Middleware
	app.Use(middleware)

	// JWT Middlware
	app.Use(jwtware.New(jwtware.Config{
		SigningKey: []byte(os.Getenv("JWT_SECRET")),
	}))

	app.Use(isAdmin)

	// Book routes
	// Group routes under /book
	bookGroup := app.Group("/books")
	bookGroup.Use(isAdmin)

	bookGroup.Get("/", getBooks)
	bookGroup.Get("/:id", getBook)
	bookGroup.Post("/", createBook)
	bookGroup.Put("/:id", updateBook)
	bookGroup.Delete("/:id", deleteBook)

	// Upload route
	app.Post("/upload", uploadFile)

	// Render HTML file route
	app.Get("/test-html", getHTML)

	// Get env route
	app.Get("/config/:env", getEnv)

	app.Listen(":8080")
}
