package v1

import (
	"github.com/go-chi/chi/v5"
	"github.com/muratovdias/url-shortner/internal/service"
)

type Router struct {
	path    string
	service *service.Service
}

func New(path string, service *service.Service) *Router {
	return &Router{
		path:    path,
		service: service,
	}
}

func (rout *Router) Routes() chi.Router {
	router := chi.NewRouter()
	router.Post("/shortner", rout.shortener)
	return router
}

func (rout *Router) Path() string {
	return rout.path
}
