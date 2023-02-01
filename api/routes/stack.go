package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/robertocamp/gd5gofcd/api/handlers"
)

// BookRouter is the router for the GoFiber app


func StackRouter(app fiber.Router) {
	app.Get("/stack", handlers.GetStack)
}