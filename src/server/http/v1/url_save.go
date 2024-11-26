package v1

import (
	"errors"
	"github.com/go-chi/render"
	"github.com/muratovdias/url-shortner/src/models"
	"github.com/muratovdias/url-shortner/src/service/shortner"
	"net/http"
	"time"
)

type UrlShortenerRequest struct {
	Url string `json:"url"`
}

type UrlShortenerResponse struct {
	Alias      string    `json:"alias"`
	ExpireTime time.Time `json:"expire_time"`
}

// @Router			/api/v1/shortener [post]
// @Summary		Создание короткой ссылки
// @Description	Метод принимает URL и возвращает короткую ссылку.
// @Tags			Короткие ссылки
// @Accept			json
// @Produce		json
// @Param			request body UrlShortenerRequest true "Данные для сокращения URL"
// @Success		201 {object} UrlShortenerResponse "Успешный ответ с созданной короткой ссылкой"
// @Failure		400 {string} string "Некорректный запрос или данные"
// @Failure		500 {string} string "Внутренняя ошибка сервера"
func (rout *Router) shortener(w http.ResponseWriter, r *http.Request) {
	var request UrlShortenerRequest
	if err := render.DecodeJSON(r.Body, &request); err != nil {
		http.Error(w, "invalid request payload", http.StatusBadRequest)
		return
	}

	response, err := rout.service.UrlShortener.Save(r.Context(), request.Url)
	if err != nil {
		if errors.Is(err, shortner.ErrInvalidUrl) || errors.Is(errors.Unwrap(err), models.ErrAlreadyExists) {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	render.Status(r, http.StatusCreated)
	render.JSON(w, r,
		UrlShortenerResponse{
			Alias:      response.Alias,
			ExpireTime: response.ExpireTime,
		})
}
