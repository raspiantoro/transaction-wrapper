package appcontext

import (
	"github.com/raspiantoro/transaction-wrapper/internal/app/driver"
	"github.com/raspiantoro/transaction-wrapper/internal/app/handler"
	"github.com/raspiantoro/transaction-wrapper/internal/app/repository"
	"github.com/raspiantoro/transaction-wrapper/internal/app/router"
	"github.com/raspiantoro/transaction-wrapper/internal/app/service"
	"github.com/spf13/viper"
)

type AppContext struct{}

func New() *AppContext {
	actx := &AppContext{}
	return actx
}

func (actx *AppContext) GetDB() (db *driver.Database, err error) {
	dbConfig := driver.DBConfig{
		Host:     viper.GetString("database.host"),
		Port:     viper.GetString("database.port"),
		Name:     viper.GetString("database.name"),
		UserName: viper.GetString("database.username"),
		Password: viper.GetString("database.password"),
	}

	db, err = driver.NewDatabase(dbConfig)

	return
}

func (actx *AppContext) GetService(repo repository.Repository) *service.Service {

	svc := service.NewService()

	personService := service.NewPersonService(repo)
	svc.SetPersonService(personService)

	return svc
}

func (actx *AppContext) GetHandler(svc *service.Service) *handler.Handler {
	h := handler.New(svc)
	return h
}

func (actx *AppContext) GetRouter(h *handler.Handler) *router.Router {
	r := router.New(h)
	return r
}

func (actx *AppContext) GetRepository(db *driver.Database) repository.Repository {
	repoOption := repository.RepositoryOption{
		DB: db,
	}

	userRepo := repository.NewUserRepository(repoOption)
	profileRepo := repository.NewProfileRepository(repoOption)

	repo := repository.Repository{
		DB:      db,
		User:    userRepo,
		Profile: profileRepo,
	}

	return repo
}
