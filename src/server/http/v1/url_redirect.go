package v1

import (
	"database/sql"
	"errors"
	"github.com/go-chi/chi/v5"
	"github.com/muratovdias/url-shortner/src/models"
	"net/http"
	"strings"
)

// @Router			/api/v1/{link} [get]
// @Summary		Перенаправление на оригинальный URL
// @Description	Метод перенаправляет пользователя на оригинальный URL, связанный с указанной короткой ссылкой.
// @Tags			Короткие ссылки
// @Param			link path string true "Алиас короткой ссылки"
// @Success		302 "Успешное перенаправление"
// @Failure		400 {string} string "Короткая ссылка истекла или недействительна"
// @Failure		404 {string} string "Оригинальный URL не найден"
// @Failure		500 {string} string "Внутренняя ошибка сервера"
func (rout *Router) redirectToOriginal(w http.ResponseWriter, r *http.Request) {
	alias := chi.URLParam(r, "link")
	if strings.TrimSpace(alias) == "" {
		http.Error(w, "missing link parameter", http.StatusBadRequest)
		return
	}

	ctx := r.Context()
	originalUrl, err := rout.service.UrlShortener.Redirect(ctx, alias)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			http.Error(w, err.Error(), http.StatusNotFound)
		}
		if errors.Is(err, models.ErrExpired) {
			http.Error(w, err.Error(), http.StatusBadRequest)
		}
		http.Error(w, "failed to fetch original URL: "+err.Error(), http.StatusInternalServerError)
		return
	}

	http.Redirect(w, r, originalUrl, http.StatusMovedPermanently)
}
