package v1

import (
	"net/http"
)

type UrlListResponse struct {
	Urls []string `json:"urls"`
}

func (rout *Router) urlsList(w http.ResponseWriter, r *http.Request) {
	userId := r.Context().Value(userID) // предположим, что userID передан в контексте

	if userId == nil {
		http.Error(w, "unauthorized", http.StatusUnauthorized)
		return
	}

	//links, err := rout.service.GetUrlsList(r.Context(), userId)
	//if err != nil {
	//	http.Error(w, "failed to retrieve links", http.StatusInternalServerError)
	//	return
	//}
	//
	//render.Status(r, http.StatusCreated)
	//render.JSON(w, r, response)
}
