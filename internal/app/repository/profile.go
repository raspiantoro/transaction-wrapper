package repository

import (
	"context"

	"github.com/raspiantoro/transaction-wrapper/internal/app/model"
)

type profileRepository struct {
	RepositoryOption
}

func NewProfileRepository(opt RepositoryOption) ProfileRepository {
	r := &profileRepository{
		RepositoryOption: opt,
	}

	return r
}

func (p *profileRepository) CreateProfile(ctx context.Context, profile *model.Profile, opts ...RepositoryFnOptions) (err error) {
	m := &MethodOption{
		RepositoryOption: p.RepositoryOption,
	}

	setMethodOption(m, opts...)

	err = m.getDB().Create(profile).Error

	return
}
