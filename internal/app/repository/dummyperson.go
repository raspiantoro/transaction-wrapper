package repository

import "github.com/raspiantoro/transaction-wrapper/internal/app/model"

type DummyPersonRepository struct {
}

func NewDummyPersonRepository() PersonRepository {
	return new(DummyPersonRepository)
}

func (d *DummyPersonRepository) GetPerson() (m []model.Person, err error) {
	return
}
