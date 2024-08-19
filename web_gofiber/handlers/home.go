package handlers

import "github.com/gofiber/fiber/v2"

func HomePage(c *fiber.Ctx) error {
	return c.SendString("Hello, Fiber App")
}
