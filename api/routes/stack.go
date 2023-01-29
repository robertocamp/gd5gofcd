package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/robertocamp/gd5gofcd/api/handlers"
	"github.com/robertocamp/gd5gofcd/pkg/book"
	"github.com/robertocamp/gd5gofcd/pkg/stack"
)

// BookRouter is the router for the GoFiber app


func StackRouter(app fiber.Router, service stack.Service) {
	app.Get("/stack", handlers.GetStack(service))
}