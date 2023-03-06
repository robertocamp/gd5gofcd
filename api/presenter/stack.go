package presenter

import (
	"github.com/robertocamp/gd5gofcd/pkg/entities"
	"github.com/gofiber/fiber/v2"
)

// Stack is the presenter object which will be passed in the response by Handler
type Stack struct {
	Items          []entities.Items `json:"items"`
}

func StackSuccessResponse(data entities.Stack) *fiber.Map {
	stack :=Stack{
		Items:							data.Items,
	}
	return &fiber.Map{
		"status": true,
		"data":   stack,
		"error":  nil,
	}
}

// StacksSuccessResponse is the list SuccessResponse that will be passed in the response by Handler
func StacksSuccessResponse(data *[]Stack) *fiber.Map {
	return &fiber.Map{
		"status": true,
		"data":   data,
		"error":  nil,
	}
}

// StackErrorResponse is the ErrorResponse that will be passed in the response by Handler
func StackErrorResponse(err error) *fiber.Map {
	return &fiber.Map{
		"status": false,
		"data":   "",
		"error":  err.Error(),
	}
}