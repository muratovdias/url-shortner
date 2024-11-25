package v1

import (
	"database/sql"
	"errors"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/render"
	"net/http"
	"time"
)

type urlStatsResponse struct {
	Clicks     int       `json:"clicks"`
	LastAccess time.Time `json:"last_access"`
}

func (rout *Router) urlStats(w http.ResponseWriter, r *http.Request) {
	alias := chi.URLParam(r, "link")
	if alias == "" {
		http.Error(w, "missing link parameter", http.StatusBadRequest)
		return
	}

	stats, err := rout.service.UrlShortener.Stats(r.Context(), alias)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			http.Error(w, err.Error(), http.StatusNotFound)
		}
		http.Error(w, "failed to fetch stats: "+err.Error(), http.StatusInternalServerError)
		return
	}

	render.Status(r, http.StatusOK)
	render.JSON(w, r, urlStatsResponse{
		Clicks:     stats.Clicks,
		LastAccess: stats.LastAccessTime,
	})
}
