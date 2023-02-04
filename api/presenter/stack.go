package presenter

import (
	"github.com/robertocamp/gd5gofcd/pkg/entities"
	"github.com/gofiber/fiber/v2"
)

// Stack is the presenter object which will be passed in the response by Handler


// StackSuccessResponse is the singular SuccessResponse that will be passed in the response by
//Handler
func StackSuccessResponse(data *entities.Stack) *fiber.Map {
	stack := entities.Stack{
		Items:     			data.Items,
		HasMore:  			data.HasMore,
		QuotaMax: 			data.QuotaMax,
		QuotaRemaining: data.QuotaRemaining,
	}
	return &fiber.Map{
		"status": true,
		"data":   stack,
		"error":  nil,
	}
}

// StacksSuccessResponse is the list SuccessResponse that will be passed in the response by Handler
func StacksSuccessResponse(data *[]entities.Stack) *fiber.Map {
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