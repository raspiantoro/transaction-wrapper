package service

import (
	"context"

	"github.com/raspiantoro/transaction-wrapper/internal/app/payload"
	"github.com/raspiantoro/transaction-wrapper/internal/app/repository"
)

type personService struct {
	ServiceOptions
}

func NewPersonService(repository repository.Repository) (s PersonService) {
	so := ServiceOptions{
		Repository: repository,
	}

	s = &personService{
		ServiceOptions: so,
	}
	return
}

func (ps *personService) GetPerson(ctx context.Context, ID string) (resp payload.GetPersonResponses, err error) {
	person, err := ps.Repository.User.GetUser(ctx, ID)
	if err != nil {
		return
	}

	resp = payload.GetPersonResponses{
		ID:        person.ID,
		UserName:  person.UserName,
		FirstName: person.Profile.FirstName,
		LastName:  person.Profile.LastName,
		Age:       person.Profile.Age,
	}

	return
}

func (ps *personService) GetPersons(ctx context.Context) (resp payload.GetPersonsResponses, err error) {

	persons, err := ps.Repository.User.GetUsers(ctx)
	if err != nil {
		return
	}

	for _, person := range persons {
		resp.Users = append(resp.Users, payload.GetPersonResponses{
			ID:        person.ID,
			UserName:  person.UserName,
			FirstName: person.Profile.FirstName,
			LastName:  person.Profile.LastName,
			Age:       person.Profile.Age,
		})
	}

	return
}
