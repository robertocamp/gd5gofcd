package stack

import (
	"github.com/robertocamp/gd5gofcd/api/presenter"
	"github.com/robertocamp/gd5gofcd/pkg/entities"
)

//Service is an interface from which our api module can access our stack API data
type Service interface {
	GetFullStack() (*[]presenter.Stack, error)
}

// type service struct {
// 	repository Repository
// }

// //NewService is used to create a single instance of the service
// func NewService(r Repository) Service {
// 	return &service{
// 		repository: r,
// 	}
// }

