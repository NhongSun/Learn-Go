package main

import (
	"fmt"
	"os"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v4"
)

type User struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

var user = User{
	Email:    "user@example.com",
	Password: "1234",
}

func login(c *fiber.Ctx) error {
	newUser := new(User)

	if err := c.BodyParser(newUser); err != nil {
		return c.Status(fiber.StatusBadRequest).SendString(err.Error())
	}

	if newUser.Email != user.Email || newUser.Password != user.Password {
		return c.Status(fiber.StatusUnauthorized).SendString("Invalid email or password")
		// or return fiber.ErrUnauthorized
	}

	// Create token
	token := jwt.New(jwt.SigningMethodHS256)

	// Set claims : store token data
	claims := token.Claims.(jwt.MapClaims)
	claims["email"] = user.Email
	claims["role"] = "admin"
	claims["exp"] = time.Now().Add(time.Hour * 72).Unix()

	// Generate encoded token and send it as response
	t, err := token.SignedString([]byte(os.Getenv("JWT_SECRET")))
	if err != nil {
		fmt.Println(err)
		return c.SendStatus(fiber.StatusInternalServerError)
	}

	return c.JSON(fiber.Map{
		"message": "Login success",
		"token":   t,
	})
}

func isAdmin(c *fiber.Ctx) error {
	user := c.Locals("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)

	if claims["role"] == "admin" {
		return fiber.ErrUnauthorized
	}

	return c.Next()
}

func middleware(c *fiber.Ctx) error {
	start := time.Now()

	fmt.Printf("URL = %s\n", c.OriginalURL())
	fmt.Printf("Method = %s\n", c.Method())
	fmt.Printf("Time = %s\n", start)
	fmt.Println("")

	return c.Next()
}
