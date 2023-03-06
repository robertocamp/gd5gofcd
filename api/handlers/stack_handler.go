package handlers

import (
	"net/http"
	"github.com/gofiber/fiber/v2"
	"github.com/robertocamp/gd5gofcd/api/presenter"
	"github.com/robertocamp/gd5gofcd/pkg/stack"
)


// GetBooks is handler/controller which retrieves the stackoverflow API
func GetStack(service stack.Service) fiber.Handler {
	return func(c *fiber.Ctx) error {
		api, err := service.NewStack()
		if err != nil {
			c.Status(http.StatusInternalServerError)
			return c.JSON(presenter.StackErrorResponse(err))
		}
		return c.JSON(presenter.StackSuccessResponse(api))
	}
}