package service

import (
	"github.com/raspiantoro/transaction-wrapper/internal/app/repository"
)

type ServiceOptions struct {
	Repository repository.Repository
}

type Service struct {
	Person PersonService
}

func NewService() *Service {
	s := new(Service)
	return s
}

func (s *Service) SetPersonService(ps PersonService) {
	s.Person = ps
}
