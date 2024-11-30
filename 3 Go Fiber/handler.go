package main

import (
	"os"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

func getBooks(c *fiber.Ctx) error {
	// request, response, header, ... is in fiber context

	return c.JSON(fiber.Map{
		"count": len(books),
		"data":  books,
	})
}

func getBook(c *fiber.Ctx) error {
	// get bookId in param and convert to string
	bookId, err := strconv.Atoi(c.Params("id"))

	if err != nil {
		return c.Status(fiber.StatusBadRequest).SendString(err.Error())
	}

	// for index, book := range books {
	// -> index is not used so just _
	for _, book := range books {
		if book.ID == bookId {
			return c.JSON(book)
		}
	}

	// return c.SendStatus(fiber.StatusNotFound)
	return c.Status(fiber.StatusNotFound).SendString("Book not found!")
}

func createBook(c *fiber.Ctx) error {
	book := new(Book)

	if err := c.BodyParser(book); err != nil {
		return c.Status(fiber.StatusBadRequest).SendString(err.Error())
	}

	books = append(books, *book)
	return c.JSON(book)
}

func updateBook(c *fiber.Ctx) error {
	bookId, err := strconv.Atoi(c.Params("id"))

	if err != nil {
		return c.Status(fiber.StatusBadRequest).SendString(err.Error())
	}

	bookUpdate := new(Book)

	if err := c.BodyParser(bookUpdate); err != nil {
		return c.Status(fiber.StatusBadRequest).SendString(err.Error())
	}

	for i, book := range books {
		if book.ID == bookId {
			books[i].Title = bookUpdate.Title
			books[i].Author = bookUpdate.Author

			return c.JSON(books[i])
		}
	}

	return c.Status(fiber.StatusNotFound).SendString("Book not found!")
}

func deleteBook(c *fiber.Ctx) error {
	bookId, err := strconv.Atoi(c.Params("id"))

	if err != nil {
		return c.Status(fiber.StatusBadRequest).SendString(err.Error())
	}

	for i, book := range books {
		if book.ID == bookId {
			// ... -> spread operator like js
			books = append(books[:i], books[i+1:]...)

			return c.SendStatus(fiber.StatusNoContent)
		}
	}

	return c.Status(fiber.StatusNotFound).SendString("Book not found!")
}

func uploadFile(c *fiber.Ctx) error {
	file, err := c.FormFile("image") // accept image file

	if err != nil {
		return c.Status(fiber.StatusBadRequest).SendString(err.Error())
	}

	err = c.SaveFile(file, "./uploads/"+file.Filename)

	if err != nil {
		return c.Status(fiber.StatusInternalServerError).SendString(err.Error())
	}

	return c.SendString("File upload complete!")
}

func getHTML(c *fiber.Ctx) error {
	return c.Render("index", fiber.Map{
		"Name": "Jeff",
	})
}

func getEnv(c *fiber.Ctx) error {
	envName := c.Params("env")
	secret := os.Getenv(envName)

	if secret == "" {
		return c.Status(fiber.StatusBadRequest).SendString("ENV not found")
	}

	return c.JSON(fiber.Map{
		"SECRET": secret,
		"a":      "a",
	})
}
