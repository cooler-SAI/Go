package handlers

import "github.com/gofiber/fiber/v2"

func BlogPage(c *fiber.Ctx) error {
	return c.SendString("This is the Blog page.")
}
