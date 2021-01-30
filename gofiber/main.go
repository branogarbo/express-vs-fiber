package main

import (
	"log"

	"github.com/gofiber/fiber/v2"
)

func main() {
	server := fiber.New()

	server.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("THIS A TEST")
	})

	log.Fatal(server.Listen(":3001"))
}
