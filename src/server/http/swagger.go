package http

import (
	"github.com/go-chi/chi/v5"
	_ "github.com/muratovdias/url-shortner/swagger"
	httpSwagger "github.com/swaggo/http-swagger"
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

	r.Get("/*", httpSwagger.WrapHandler)

	return r
}

func (sr *SwaggerResource) Path() string {
	return sr.mountPath
}
