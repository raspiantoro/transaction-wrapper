package repository

import (
	"github.com/raspiantoro/transaction-wrapper/internal/app/driver"
	"gorm.io/gorm"
)

type RepositoryOption struct {
	DB *driver.Database
}

type MethodOption struct {
	RepositoryOption
	opHandler *gorm.DB
}

func (m *MethodOption) getDB() *gorm.DB {
	if m.opHandler == nil {
		m.opHandler = m.DB.Instance
	}
	return m.opHandler
}

func setMethodOption(m *MethodOption, opts ...RepositoryFnOptions) {
	for _, opt := range opts {
		opt(m)
	}
}

type RepositoryFnOptions func(m *MethodOption)

func WithDBInstance(db *gorm.DB) RepositoryFnOptions {
	return func(m *MethodOption) {
		m.opHandler = db
	}
}

type Repository struct {
	DB      *driver.Database
	User    UserRepository
	Profile ProfileRepository
}

func (r *Repository) WithTransactions(fn func(tx *gorm.DB) error) (err error) {
	err = r.DB.Instance.Transaction(func(tx *gorm.DB) error {
		return fn(tx)
	})

	return
}
