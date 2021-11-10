package routers

import (
	"net/http"

	"github.com/go-chi/chi/v5"
)

func ItemsRouter() chi.Router {
	r := chi.NewRouter()

	r.Get("/", func(rw http.ResponseWriter, r *http.Request) {
		rw.Write([]byte("hello items"))
	})

	return r
}

