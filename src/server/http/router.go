package http

import (
	"github.com/go-chi/chi/v5"
	"github.com/muratovdias/url-shortner/src/databases/drivers"
	v1 "github.com/muratovdias/url-shortner/src/server/http/v1"
	"github.com/muratovdias/url-shortner/src/service"
)

type Router interface {
	Routes() chi.Router
	Path() string
}

type RouterImpl struct {
	service *service.Service
	ds      drivers.Base
}

func NewRouterImpl(service *service.Service, ds drivers.Base) *RouterImpl {
	return &RouterImpl{
		service: service,
		ds:      ds,
	}
}

func (r *RouterImpl) Routes() chi.Router {
	router := chi.NewRouter()

	for _, rout := range []Router{
		v1.New("/api/v1", r.service),
		NewHealthResource("/health", r.ds),
		NewSwaggerResource("/swagger", "/swagger"),
	} {
		router.Mount(rout.Path(), rout.Routes())
	}

	return router
}
