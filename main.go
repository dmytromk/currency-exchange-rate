package main

import (
	"currency_exchange_rate/internal/api"
	"currency_exchange_rate/internal/database"
	"database/sql"
	"log"

	"fmt"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	_ "github.com/lib/pq"
)

type Env struct {
	users database.UserModel
}

func main() {
	r := chi.NewRouter()

	db, err := sql.Open("postgres", "user=postgres dbname=example sslmode=disable")
	if err != nil {
		log.Fatal(err)
	}

	env := &Env{
		users: database.UserModel{DB: db},
	}

	r.Use(middleware.Logger)

	r.Get("/rate", func(w http.ResponseWriter, r *http.Request) {
		rate, err := api.GetCurrentNBURate()

		if err != nil {
			http.Error(w, http.StatusText(400), 400)
			return
		}

		w.Write([]byte(fmt.Sprintf("%f", rate)))
	})

	r.Post("/subscribe", func(w http.ResponseWriter, r *http.Request) {
		email := "test"

		if email == "" {
			http.Error(w, http.StatusText(400), 400)
			return
		}

		err = database.AddUser(database.User{Email: email}, env.users)
		if err != nil {
			http.Error(w, http.StatusText(400), 400)
			return
		}
	})

	http.ListenAndServe(":3000", r)
}
