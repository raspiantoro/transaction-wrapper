package service

import (
	"context"

	"github.com/raspiantoro/transaction-wrapper/internal/app/payload"
)

type PersonService interface {
	GetPerson(ctx context.Context, ID string) (payload.GetPersonResponses, error)
	GetPersons(ctx context.Context) (payload.GetPersonsResponses, error)
	CreatePerson(ctx context.Context, person payload.CreatePersonRequests) (err error)
	CreatePersonTx(ctx context.Context, person payload.CreatePersonRequests) (err error)
}
