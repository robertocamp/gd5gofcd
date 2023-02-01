package handlers

import "github.com/gofiber/fiber/v2"


// smoke test Hello endpoint
func GetStack(c *fiber.Ctx)  error {
	return c.SendString("Hello, World 👋!")
}

