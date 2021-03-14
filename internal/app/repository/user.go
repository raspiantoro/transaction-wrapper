package repository

import (
	"context"

	"github.com/raspiantoro/transaction-wrapper/internal/app/model"
)

type userRepository struct {
	RepositoryOption
}

func NewUserRepository(opt RepositoryOption) UserRepository {
	r := &userRepository{
		RepositoryOption: opt,
	}

	return r
}

func (u *userRepository) GetUser(ctx context.Context, ID string) (user model.User, err error) {
	err = u.DB.Instance.Preload("Profile").First(&user, "id", ID).Error
	return
}

func (u *userRepository) GetUsers(ctx context.Context) (users []model.User, err error) {
	err = u.DB.Instance.Preload("Profile").Find(&users).Debug().Error
	return
}

func (u *userRepository) CreateUser(ctx context.Context, user model.User) (err error) {
	return
}
