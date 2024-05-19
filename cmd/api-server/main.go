package main

import (
	"currency_exchange_rate/internal/api"
	"currency_exchange_rate/internal/database"
	"github.com/joho/godotenv"
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	_ "github.com/lib/pq"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	db, err := database.OpenConnection()
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

	// test endpoint for sending emails
	// should be removed from public later
	r.Post("/test_email", func(w http.ResponseWriter, r *http.Request) { env.SendEmails() })

	http.ListenAndServe(":3000", r)
}
