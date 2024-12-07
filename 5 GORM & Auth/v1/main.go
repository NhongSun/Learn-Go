package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt"
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

var DB *gorm.DB

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
	db.AutoMigrate(&Book{}, &User{})
	DB = db

	app := fiber.New()

	app.Use("/books", authRequired)

	app.Get("/books", getBooksHandler)
	app.Get("/books/:id", getBookHandler)
	app.Post("/books", createBookHandler)
	app.Put("/books/:id", updateBookHandler)
	app.Delete("/books/:id", deleteBookHandler)

	app.Post("/register", registerHandler)
	app.Post("/login", loginHandler)

	app.Listen(":8080")
}

func getBooksHandler(c *fiber.Ctx) error {
	books := getBooks(DB)

	return c.JSON(fiber.Map{
		"count": len(books),
		"data":  books,
	})
}

func getBookHandler(c *fiber.Ctx) error {
	bookId, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return c.SendStatus(fiber.StatusBadRequest)
	}

	book, err := getBook(DB, uint(bookId))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Failed to retrieve book",
		})
	}

	return c.JSON(book)
}

func createBookHandler(c *fiber.Ctx) error {
	book := new(Book)

	err := c.BodyParser(book)
	if err != nil {
		return c.SendStatus(fiber.StatusBadRequest)
	}

	err = createBook(DB, book)
	if err != nil {
		return c.SendStatus(fiber.StatusBadRequest)
	}

	return c.JSON(fiber.Map{
		"message": "Create book successful",
	})
}

func updateBookHandler(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return c.SendStatus(fiber.StatusBadRequest)
	}

	book := new(Book)
	if err := c.BodyParser(book); err != nil {
		return c.SendStatus(fiber.StatusBadRequest)
	}

	book.ID = uint(id)

	err = updateBook(DB, book)

	if err != nil {
		return c.SendStatus(fiber.StatusBadRequest)
	}

	return c.JSON(fiber.Map{
		"message": "Update book successful",
	})
}

func deleteBookHandler(c *fiber.Ctx) error {
	bookId, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return c.SendStatus(fiber.StatusBadRequest)
	}

	err = deleteBook(DB, uint(bookId))
	if err != nil {
		return c.SendStatus(fiber.StatusBadRequest)
	}

	return c.JSON(fiber.Map{
		"message": "Delete book successful",
	})
}

func registerHandler(c *fiber.Ctx) error {
	user := new(User)

	if err := c.BodyParser(user); err != nil {
		return c.SendStatus(fiber.StatusBadRequest)
	}

	err := createUser(DB, user)

	if err != nil {
		return c.SendStatus(fiber.StatusBadRequest)
	}

	return c.JSON(fiber.Map{
		"message": "Create user successful",
	})
}

func loginHandler(c *fiber.Ctx) error {
	user := new(User)

	if err := c.BodyParser(user); err != nil {
		return c.SendStatus(fiber.StatusBadRequest)
	}

	token, err := login(DB, user)

	if err != nil {
		fmt.Println(err)
		return c.SendStatus(fiber.StatusUnauthorized)
	}

	// Set cookie
	c.Cookie(&fiber.Cookie{
		Name:     "jwt",
		Value:    token,
		Expires:  time.Now().Add(time.Hour * 72),
		HTTPOnly: true,
	})

	return c.JSON(fiber.Map{
		"token": token,
	})
}

func authRequired(c *fiber.Ctx) error {
	cookie := c.Cookies("jwt")
	jwtSecretKey := os.Getenv("JWT_SECRET_KEY")

	token, err := jwt.ParseWithClaims(cookie, jwt.MapClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(jwtSecretKey), nil
	})

	if err != nil || !token.Valid {
		return c.SendStatus(fiber.StatusUnauthorized)
	}

	claim := token.Claims.(jwt.MapClaims)

	fmt.Println(claim["user_id"])

	return c.Next()
}
