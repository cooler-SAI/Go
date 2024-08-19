package handlers

import "github.com/gofiber/fiber/v2"

func ContactPage(c *fiber.Ctx) error {
	return c.SendString("This is the Contact page.")
}
