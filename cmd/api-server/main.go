package main

import (
	"currency_exchange_rate/internal/api"
	"currency_exchange_rate/internal/database"
	"database/sql"
	"github.com/go-chi/chi/v5"
	"log"

	"net/http"

	"github.com/go-chi/chi/v5/middleware"
	_ "github.com/lib/pq"
)

func main() {
	db, err := sql.Open("postgres", "user=postgres dbname=example sslmode=disable")
	if err != nil {
		log.Fatal(err)
	}

	env := &api.Env{
		Users: database.UserModel{DB: db},
	}

	r := chi.NewRouter()
	r.Use(middleware.Logger)

	r.Get("/rate", env.GetCurrentRate)
	r.Post("/subscribe", env.AddUser)

	http.ListenAndServe(":3000", r)
}
