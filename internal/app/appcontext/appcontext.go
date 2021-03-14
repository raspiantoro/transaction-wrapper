package appcontext

import (
	"github.com/raspiantoro/transaction-wrapper/internal/app/handler"
	"github.com/raspiantoro/transaction-wrapper/internal/app/repository"
	"github.com/raspiantoro/transaction-wrapper/internal/app/router"
	"github.com/raspiantoro/transaction-wrapper/internal/app/service"
)

type AppContext struct{}

func New() *AppContext {
	actx := &AppContext{}
	return actx
}

func (actx *AppContext) GetService(repo repository.PersonRepository) *service.Service {

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

func (actx *AppContext) GetRepository() repository.PersonRepository {
	repo := repository.NewDummyPersonRepository()
	return repo
}
