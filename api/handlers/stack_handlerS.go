
package handlers

import (
	"github.com/robertocamp/gd5gofcd/api/presenter"
	"github.com/robertocamp/gd5gofcd/pkg/stack"
	"github.com/robertocamp/gd5gofcd/pkg/entities"
	"github.com/gofiber/fiber/v2"
	"github.com/pkg/errors"
	"net/http"
)

// GetBooks is handler/controller which gets the full stack API
func GetStackS(service stack.Service) fiber.Handler {
	return func(c *fiber.Ctx) error {
		fetched, err := service.StackAPI()
		if err != nil {
			c.Status(http.StatusInternalServerError)
			return c.JSON(presenter.BookErrorResponse(err))
		}
		return c.JSON(presenter.StackSuccessResponse(fetched))
	}
}

// // GetBooks is handler/controller which lists all Books from the BookShop
// func GetBooks(service book.Service) fiber.Handler {
// 	return func(c *fiber.Ctx) error {
// 		fetched, err := service.FetchBooks()
// 		if err != nil {
// 			c.Status(http.StatusInternalServerError)
// 			return c.JSON(presenter.BookErrorResponse(err))
// 		}
// 		return c.JSON(presenter.BooksSuccessResponse(fetched))
// 	}
// }