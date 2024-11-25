package v1

import (
	"database/sql"
	"errors"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/render"
	"net/http"
)

// @Router			/api/v1/{link} [delete]
// @Summary	удаления короткой ссылки.
// @Description	Запрос на удаление короткой ссылки.
// @Tags			Удаление
// @Param link path string true
// @Success		204
// @Failure		400
// @Failure		500
func (rout *Router) deleteShortLink(w http.ResponseWriter, r *http.Request) {
	link := chi.URLParam(r, "link")

	err := rout.service.UrlShortener.Delete(r.Context(), link)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			http.Error(w, "link not found", http.StatusNotFound)
			return
		}

		http.Error(w, "failed to delete the short link", http.StatusInternalServerError)
		return
	}

	render.Status(r, http.StatusNoContent)
}
