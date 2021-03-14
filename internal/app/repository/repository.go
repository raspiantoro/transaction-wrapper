package repository

import (
	"github.com/raspiantoro/transaction-wrapper/internal/app/driver"
	"gorm.io/gorm"
)

type RepositoryOption struct {
	DB *driver.Database
}

type OperationHandler struct {
	opHandler *gorm.DB
}

type Repository struct {
	User    UserRepository
	Profile ProfileRepository
}
