package main

import (
	"fmt"
	"sync"

	"github.com/gofiber/fiber/v2"
)

type Message struct {
	Data string `json:"data"`
}

type PubSub struct {
	subs []chan Message // slice of channel of Message
	mu   sync.Mutex
}

// method to publish a message of PubSub
func (ps *PubSub) Subscribe() chan Message {
	ps.mu.Lock()
	defer ps.mu.Unlock()

	ch := make(chan Message, 1)
	ps.subs = append(ps.subs, ch)
	return ch
}

// method to publish a message of PubSub
func (ps *PubSub) Publish(msg *Message) {
	ps.mu.Lock()
	defer ps.mu.Unlock()

	for _, sub := range ps.subs {
		sub <- *msg
	}
}

func (ps *PubSub) Unsubscribe(ch chan Message) {
	ps.mu.Lock()
	defer ps.mu.Unlock()

	for i, sub := range ps.subs {
		if sub == ch {
			ps.subs = append(ps.subs[:i], ps.subs[i+1:]...)
			close(ch)
			break
		}
	}
}

func main() {
	app := fiber.New()

	pubsub := &PubSub{}

	app.Post("/publish", func(c *fiber.Ctx) error {
		message := new(Message)

		if err := c.BodyParser(message); err != nil {
			return c.SendStatus(fiber.StatusBadRequest)
		}

		pubsub.Publish(message)
		return c.JSON(fiber.Map{"message": "Add to subscriber"})
	})

	sub := pubsub.Subscribe()
	go func() {
		for msg := range sub {
			fmt.Println("Receive message", msg)
		}
	}()

	app.Listen(":8080")
}
