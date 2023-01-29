package handlers

import "github.com/gofiber/fiber/v2"


func GetStack(c *fiber.Ctx)  error {
	return c.SendString("Hello, World 👋!")
}


// func GetStack(service stack.Service) fiber.Handler {
// 	return func(c *fiber.Ctx) error {
// 		fetched, err := service.Get
// 		if err != nil {
// 			c.Status(http.StatusInternalServerError)
// 			return c.JSON(presenter.BookErrorResponse(err))
// 		}
// 		return c.JSON(presenter.BooksSuccessResponse(fetched))
// 	}
// }