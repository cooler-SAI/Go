package handlers

import "github.com/gofiber/fiber/v2"

func AboutPage(c *fiber.Ctx) error {
	return c.SendString("This is the About page.")
}
