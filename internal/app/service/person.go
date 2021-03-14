package service

import (
	"context"

	"github.com/raspiantoro/transaction-wrapper/internal/app/repository"

	"github.com/raspiantoro/transaction-wrapper/internal/app/model"
)

type personService struct {
	ServiceOptions
}

func NewPersonService(personRepo repository.PersonRepository) (s PersonService) {
	so := ServiceOptions{
		PersonRepository: personRepo,
	}

	s = &personService{
		ServiceOptions: so,
	}
	return
}

func (ps *personService) GetPerson(ctx context.Context) (persons []model.Person, err error) {
	persons, err = ps.PersonRepository.GetPerson()
	return
}
