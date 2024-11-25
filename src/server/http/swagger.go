package http

import (
	"github.com/go-chi/chi/v5"
	"net/http"
)

type SwaggerResource struct {
	mountPath string
	basePath  string
}

func NewSwaggerResource(mountPath, basePath string) *SwaggerResource {
	return &SwaggerResource{
		mountPath: mountPath,
		basePath:  basePath,
	}
}

func (sr *SwaggerResource) Routes() chi.Router {
	r := chi.NewRouter()

	r.Route("/swagger", func(r chi.Router) {
		// Сервинг файлов в папке "./swagger"
		fs := http.StripPrefix("/swagger/", http.FileServer(http.Dir("./swagger")))
		r.Get("/*", http.StripPrefix("/swagger/", fs).ServeHTTP)
	})

	return r
}

func (sr *SwaggerResource) Path() string {
	return sr.mountPath
}
