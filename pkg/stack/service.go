package stack

import (
	"github.com/robertocamp/gd5gofcd/api/presenter"
	"github.com/robertocamp/gd5gofcd/pkg/entities"
)


type Service interface {}

type service struct {}

//NewService is used to create a single instance of the service
func NewService(s string) Service {
	return &service{}
}