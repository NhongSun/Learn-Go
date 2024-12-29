package main

import (
	"nhongsun/adapters"
	"nhongsun/core"

	"github.com/gofiber/fiber/v2"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func main() {
	app := fiber.New()

	db, err := gorm.Open(sqlite.Open("orders.db"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	// secondary -> primary
	orderRepo := adapters.NewGormOrderRepository(db)           // secondary port
	orderService := core.NewOrderService(orderRepo)            // primary port
	orderHandler := adapters.NewHttpOrderHandler(orderService) // primary adapter

	app.Post("/order", orderHandler.CreateOrder)

	db.AutoMigrate(&core.Order{})

	app.Listen(":8080")
}
