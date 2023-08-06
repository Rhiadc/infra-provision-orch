package domain

import "fmt"

type Service struct {
}

func NewService() *Service {
	return &Service{}
}

func (s *Service) PushToGit() {
	fmt.Println("Pushed")
}
