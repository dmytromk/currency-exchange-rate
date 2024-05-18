package main

import (
	"currency_exchange_rate/internal/api"

	"fmt"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func main() {
	fmt.Println(api.GetCurrentNBURate())

	r := chi.NewRouter()

	r.Use(middleware.Logger)

	r.Get("/rate", func(w http.ResponseWriter, r *http.Request) {
		rate, err := api.GetCurrentNBURate()

		if err != nil {
			http.Error(w, http.StatusText(400), 400)
			return
		}

		w.Write([]byte(fmt.Sprintf("%f", rate)))
	})

	http.ListenAndServe(":3000", r)
}
