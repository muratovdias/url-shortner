package v1

import (
	"database/sql"
	"errors"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/render"
	"net/http"
	"strings"
)

// @Router			/api/v1/{link} [delete]
// @Summary	Удаления короткой ссылки.
// @Description	Запрос на удаление короткой ссылки.
// @Tags			Удаление
// @Param link path string true "Алиас короткой ссылки"
// @Success		204 "Успешное удаление"
// @Failure		400 {string} string "Некорректный алиас"
// @Failure		404 {string} string "Короткая ссылка не найдена"
// @Failure		500 {string} string "Внутренняя ошибка сервера"
func (rout *Router) deleteShortLink(w http.ResponseWriter, r *http.Request) {
	link := chi.URLParam(r, "link")

	if strings.TrimSpace(link) == "" {
		http.Error(w, "link is required", http.StatusBadRequest)
		return
	}

	if err := rout.service.UrlShortener.Delete(r.Context(), link); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			http.Error(w, "link not found", http.StatusNotFound)
			return
		}

		http.Error(w, "failed to delete the short link", http.StatusInternalServerError)
		return
	}

	render.Status(r, http.StatusOK)
}
