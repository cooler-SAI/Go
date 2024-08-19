package main

import (
	"github.com/gofiber/fiber/v2"
	"web_gofiber/handlers"
)

func main() {
	fiberApp := fiber.New()
	fiberApp.Get("/", handlers.HomePage)
	fiberApp.Get("/about", handlers.AboutPage)
	fiberApp.Get("/contact", handlers.ContactPage)
	fiberApp.Get("/blog", handlers.BlogPage)

	err := fiberApp.Listen(":3000")
	if err != nil {
		return
	}
}
