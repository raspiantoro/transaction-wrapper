package service

import (
	"context"

	"github.com/raspiantoro/transaction-wrapper/internal/app/model"
	"github.com/raspiantoro/transaction-wrapper/internal/app/payload"
	"github.com/raspiantoro/transaction-wrapper/internal/app/repository"
	uuid "github.com/satori/go.uuid"
	"gorm.io/gorm"
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

func (ps *personService) CreatePerson(ctx context.Context, person payload.CreatePersonRequests) (err error) {

	user := &model.User{
		ID:       uuid.NewV4().String(),
		UserName: person.UserName,
		Password: person.Password,
	}

	profile := &model.Profile{
		ID:        uuid.NewV4().String(),
		FirstName: person.FirstName,
		LastName:  person.LastName,
		Age:       person.Age,
	}

	err = ps.Repository.User.CreateUser(ctx, user)
	if err != nil {
		return
	}

	// commenting the line below to throw error on CreateProfile
	// will resulting inconsistency on database
	profile.UserID = user.ID

	err = ps.Repository.Profile.CreateProfile(ctx, profile)

	return
}

func (ps *personService) CreatePersonTx(ctx context.Context, person payload.CreatePersonRequests) (err error) {

	user := &model.User{
		ID:       uuid.NewV4().String(),
		UserName: person.UserName,
		Password: person.Password,
	}

	profile := &model.Profile{
		ID:        uuid.NewV4().String(),
		FirstName: person.FirstName,
		LastName:  person.LastName,
		Age:       person.Age,
	}

	err = ps.Repository.WithTransactions(func(tx *gorm.DB) (err error) {

		err = ps.Repository.User.CreateUser(ctx, user, repository.WithDBInstance(tx))
		if err != nil {
			return
		}

		// commenting the line below to throw error on CreateProfile
		// won't resulting inconsistency on database, CreateUser will be rollback
		profile.UserID = user.ID

		err = ps.Repository.Profile.CreateProfile(ctx, profile, repository.WithDBInstance(tx))

		return
	})

	return
}
