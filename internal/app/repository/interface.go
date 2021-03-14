package repository

import (
	"github.com/raspiantoro/transaction-wrapper/internal/app/model"
)

type PersonRepository interface {
	GetPerson() ([]model.Person, error)
}
