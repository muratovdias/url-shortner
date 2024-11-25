package v1

import (
	"github.com/go-chi/render"
	"net/http"
)

type UrlShortenerRequest struct {
	Url string `json:"url"`
}

func (rout *Router) shortener(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	userId, _ := ctx.Value(userID).(string)

	if userId == "" {
		http.Error(w, "unauthorized", http.StatusUnauthorized)
		return
	}

	var request UrlShortenerRequest
	if err := render.DecodeJSON(r.Body, &request); err != nil {
		http.Error(w, "invalid request payload", http.StatusBadRequest)
		return
	}

	response, err := rout.service.UrlShortener.Save(ctx, userId, request.Url)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	render.Status(r, http.StatusCreated)
	render.JSON(w, r, response)
}
