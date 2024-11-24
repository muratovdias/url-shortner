package v1

import (
	"github.com/go-chi/render"
	"github.com/muratovdias/url-shortner/internal/models"
	"net/http"
)

func (rout *Router) shortener(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	var request models.Link
	if err := render.DecodeJSON(r.Body, &request); err != nil {
		http.Error(w, "invalid request payload", http.StatusBadRequest)
		return
	}

	response, err := rout.service.UrlShortener.Save(ctx, request.Url)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	render.Status(r, http.StatusCreated)
	render.JSON(w, r, response)
}
