package main

import (
	"log"

	"github.com/gofiber/fiber/v2"
)

type Json struct {
	Data string `json:"data"`
}

func main() {
	server := fiber.New(fiber.Config{
		Prefork: true,
	})

	server.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("THIS A TEST")
	})

	server.Get("/json", func(c *fiber.Ctx) error {
		data := Json{
			Data: "hello there",
		}

		return c.JSON(data)
	})

	log.Fatal(server.Listen(":3001"))
}
