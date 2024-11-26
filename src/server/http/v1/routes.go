package v1

import (
	"github.com/go-chi/chi/v5"
	"github.com/muratovdias/url-shortner/src/service"
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

	router.Post("/shortener", rout.shortener)
	router.Get("/shortener", rout.urlsList)
	router.Get("/{link}", rout.redirectToOriginal)
	router.Delete("/{link}", rout.deleteShortLink)
	router.Get("/stats/{link}", rout.urlStats)

	return router
}

func (rout *Router) Path() string {
	return rout.path
}
