package v1

import (
	"database/sql"
	"errors"
	"github.com/go-chi/chi/v5"
	"github.com/muratovdias/url-shortner/src/models"
	"net/http"
)

func (rout *Router) redirectToOriginal(w http.ResponseWriter, r *http.Request) {
	alias := chi.URLParam(r, "link")
	if alias == "" {
		http.Error(w, "missing link parameter", http.StatusBadRequest)
		return
	}

	ctx := r.Context()
	originalUrl, err := rout.service.UrlShortener.Redirect(ctx, alias)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			http.Error(w, err.Error(), http.StatusNoContent)
		}
		if errors.Is(err, models.ErrExpired) {
			http.Error(w, err.Error(), http.StatusBadRequest)
		}
		http.Error(w, "failed to fetch original URL: "+err.Error(), http.StatusNotFound)
		return
	}

	http.Redirect(w, r, originalUrl, http.StatusFound)
}
