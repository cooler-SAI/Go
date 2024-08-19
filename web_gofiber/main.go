package main

import "github.com/gofiber/fiber/v2"

func main() {
	fiberApp := fiber.New()
	fiberApp.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, gofiber!!!")
	})

	fiberApp.Get("/about", func(c *fiber.Ctx) error {
		return c.SendString("This is the About page.")
	})

	fiberApp.Get("/contact", func(c *fiber.Ctx) error {
		return c.SendString("This is the Contact page.")
	})

	fiberApp.Get("/blog/:id", func(c *fiber.Ctx) error {
		return c.SendString("Blog ID: " + c.Params("id"))
	})

	err := fiberApp.Listen(":3000")
	if err != nil {
		return
	}
}
