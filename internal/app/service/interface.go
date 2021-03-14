package service

import (
	"context"

	"github.com/raspiantoro/transaction-wrapper/internal/app/model"
)

type PersonService interface {
	GetPerson(ctx context.Context) ([]model.Person, error)
}
