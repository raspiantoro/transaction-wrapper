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

func (u *userRepository) GetUser(ctx context.Context, ID string, opts ...RepositoryFnOptions) (user model.User, err error) {
	m := &MethodOption{
		RepositoryOption: u.RepositoryOption,
	}

	setMethodOption(m, opts...)

	err = m.getDB().Preload("Profile").First(&user, "id", ID).Error
	return
}

func (u *userRepository) GetUsers(ctx context.Context, opts ...RepositoryFnOptions) (users []model.User, err error) {
	m := &MethodOption{
		RepositoryOption: u.RepositoryOption,
	}

	setMethodOption(m, opts...)

	err = m.getDB().Preload("Profile").Find(&users).Debug().Error
	return
}

func (u *userRepository) CreateUser(ctx context.Context, user *model.User, opts ...RepositoryFnOptions) (err error) {

	m := &MethodOption{
		RepositoryOption: u.RepositoryOption,
	}

	setMethodOption(m, opts...)

	err = m.getDB().Create(user).Error

	return
}
