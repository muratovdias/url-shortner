package v1

import (
	"context"
	"crypto/rand"
	"github.com/go-chi/chi/v5"
	"github.com/muratovdias/url-shortner/internal/service"
	"math"
	"math/big"
	"net/http"
	"strconv"
)

type ctxKey int

const (
	userID ctxKey = iota
)

type Router struct {
	path    string
	service *service.Service
}

func New(path string, service *service.Service) *Router {
	return &Router{
		path:    path,
		service: service,
	}
}

func generateID() string {
	index, _ := rand.Int(rand.Reader, big.NewInt(math.MaxInt32))

	return strconv.Itoa(int(index.Int64()))
}

func (rout *Router) Routes() chi.Router {
	router := chi.NewRouter()

	router.Use(func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			var id string

			cookie, err := r.Cookie("user_id")
			if err != nil {
				id = generateID()

				r.AddCookie(&http.Cookie{
					Name:  "user_id",
					Value: id,
				})

				ctx := context.WithValue(context.Background(), userID, id)
				r.WithContext(ctx)
			} else {
				id = cookie.Value
				ctx := context.WithValue(context.Background(), userID, id)
				r.WithContext(ctx)
			}

			next.ServeHTTP(w, r)
		})
	})

	router.Post("/shortner", rout.shortener)
	router.Get("/shortener", rout.urlsList)
	//router.Get("/{link}", rout.redirectToOriginal)
	//router.Delete("/{link}", rout.deleteShortLink)
	return router
}

func (rout *Router) Path() string {
	return rout.path
}
