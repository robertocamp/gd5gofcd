package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/robertocamp/gd5gofcd/api/handlers"
)



// basic route w/o  service
func StackRouter(app fiber.Router) {
	app.Get("/stack", handlers.GetStack)
}


// BookRouter is the router for the GoFiber app

func StackRouterS(app fiber.Router, service stack.Service) {
	app.Get("/stacks", handlers.GetStackS(service))
}

