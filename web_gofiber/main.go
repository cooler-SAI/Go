package main

import "github.com/gofiber/fiber/v2"

func main() {
	fiberApp := fiber.New()
	fiberApp.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, gofiber!!!")
	})

	err := fiberApp.Listen(":3000")
	if err != nil {
		return
	}
}
