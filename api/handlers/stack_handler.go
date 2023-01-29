package handlers

import "github.com/gofiber/fiber/v2"


func GetStack(c *fiber.Ctx)  error {
	return c.SendString("Hello, World 👋!")
}
