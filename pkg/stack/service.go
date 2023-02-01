package stack

import (
)


type Service interface {
}

type service struct {}

//NewService is used to create a single instance of the service
func NewService(s string) Service {
	return &service{}
}