package http

import (
	"github.com/go-chi/chi/v5"
	v1 "github.com/muratovdias/url-shortner/src/server/http/v1"
	"github.com/muratovdias/url-shortner/src/service"
)

type Router interface {
	Routes() chi.Router
	Path() string
}

type RouterImpl struct {
	service *service.Service
}

func NewRouterImpl(service *service.Service) *RouterImpl {
	return &RouterImpl{service: service}
}

func (r *RouterImpl) Routes() chi.Router {
	router := chi.NewRouter()

	for _, rout := range []Router{
		v1.New("/api/v1", r.service),
		NewSwaggerResource("/swagger", "/swagger"),
	} {
		router.Mount(rout.Path(), rout.Routes())
	}

	return router
}
