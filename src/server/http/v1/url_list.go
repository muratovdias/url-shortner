package v1

import (
	"github.com/go-chi/render"
	"net/http"
	"time"
)

type UrlsListResponse struct {
	Url     string    `json:"url"`
	Alias   string    `json:"alias"`
	Expires time.Time `json:"expires"`
}

func (rout *Router) urlsList(w http.ResponseWriter, r *http.Request) {
	links, err := rout.service.UrlShortener.GetUrlsList(r.Context())
	if err != nil {
		http.Error(w, "failed to retrieve links", http.StatusInternalServerError)
		return
	}

	render.Status(r, http.StatusOK)
	render.JSON(w, r, func() []UrlsListResponse {
		response := make([]UrlsListResponse, len(links))
		for _, link := range links {
			response = append(response, UrlsListResponse{
				Url:     link.Url,
				Alias:   link.Alias,
				Expires: link.ExpireTime,
			})
		}
		return response
	}())
}
