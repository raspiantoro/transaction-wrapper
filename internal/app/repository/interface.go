package repository

import (
	"context"

	"github.com/raspiantoro/transaction-wrapper/internal/app/model"
)

type PersonRepository interface {
	GetPerson() ([]model.Person, error)
}

type UserRepository interface {
	GetUser(ctx context.Context, ID string) (user model.User, err error)
	GetUsers(ctx context.Context) (users []model.User, err error)
	CreateUser(ctx context.Context, user model.User) (err error)
}

type ProfileRepository interface {
	GetProfile(ctx context.Context, ID uint64) (profile model.Profile, err error)
	GetProfiles(ctx context.Context) (profiles []model.Profile, err error)
	GetProfleByUserID(ctx context.Context, userID uint64) (profile model.Profile, err error)
	CreateProfile(ctx context.Context, profile model.Profile) (err error)
}
