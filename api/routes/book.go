package routes

import (
	"github.com/robertocamp/gd5gofcd/api/handlers"
	"github.com/robertocamp/gd5gofcd/pkg/book"
	"github.com/gofiber/fiber/v2"
)

// BookRouter is the router for the GoFiber app
func BookRouter(app fiber.Router, service book.Service) {
	app.Get("/books", handlers.GetBooks(service))
	app.Post("/books", handlers.AddBook(service))
	app.Put("/books", handlers.UpdateBook(service))
	app.Delete("/books", handlers.RemoveBook(service))
}

