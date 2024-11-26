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

// @Router			/api/v1/{link}/stats [get]
// @Summary		Получение статистики короткой ссылки
// @Description	Метод возвращает статистику по указанной короткой ссылке (количество переходов и последнее время доступа).
// @Tags			Короткие ссылки
// @Param			link path string true "Алиас короткой ссылки"
// @Success		200 {object} urlStatsResponse "Успешный ответ со статистикой"
// @Failure		400 {string} string "Некорректный алиас"
// @Failure		404 {string} string "Короткая ссылка не найдена"
// @Failure		500 {string} string "Внутренняя ошибка сервера"
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
