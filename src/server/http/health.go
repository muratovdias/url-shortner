package http

import (
	"github.com/go-chi/chi/v5"
	"github.com/muratovdias/url-shortner/src/databases/drivers"
	_ "github.com/muratovdias/url-shortner/swagger"
	"net/http"
)

type HealthCheck struct {
	path string
	ds   drivers.Base
}

func NewHealthResource(path string, ds drivers.Base) *HealthCheck {
	return &HealthCheck{
		path: path,
		ds:   ds,
	}
}

func (h *HealthCheck) Routes() chi.Router {
	r := chi.NewRouter()

	r.Get("/", h.handler)

	return r
}

func (h *HealthCheck) Path() string {
	return h.path
}

func (h *HealthCheck) handler(w http.ResponseWriter, _ *http.Request) {
	if err := h.ds.Ping(); err != nil {
		http.Error(w, err.Error(), http.StatusServiceUnavailable)
		return
	}

	w.WriteHeader(http.StatusOK)
}
