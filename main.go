package main

import (
	"fmt"
	"net/http"
	"time"

	"github.com/ben39053372/hello-gochi/routers"
	"github.com/ben39053372/hello-gochi/utils"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)


func main() {

	fmt.Println("Starting Server")

	r := chi.NewRouter()

	utils.ConnectDataBase()

	r.Use(middleware.RequestID)
	r.Use(middleware.RealIP)
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)

	r.Use(middleware.Timeout(60 * time.Second))

	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello World"))
	})

	r.Mount("/items",routers.ItemsRouter())

	http.ListenAndServe(":3000", r)
}