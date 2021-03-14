package repository

import (
	"context"

	"github.com/raspiantoro/transaction-wrapper/internal/app/model"
)

type PersonRepository interface {
	GetPerson() ([]model.Person, error)
}

type UserRepository interface {
	GetUser(ctx context.Context, ID string, opts ...RepositoryFnOptions) (user model.User, err error)
	GetUsers(ctx context.Context, opts ...RepositoryFnOptions) (users []model.User, err error)
	CreateUser(ctx context.Context, user *model.User, opts ...RepositoryFnOptions) (err error)
}

type ProfileRepository interface {
	CreateProfile(ctx context.Context, profile *model.Profile, opts ...RepositoryFnOptions) (err error)
}
