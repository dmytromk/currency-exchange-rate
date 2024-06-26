package api

import (
	"currency_exchange_rate/internal/database"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/lib/pq"
	"log"
	"net/http"
	"net/mail"
)

type Env struct {
	Users database.UserModel
}

func (env *Env) SendEmails() {
	users, err := env.Users.GetAllUsers()
	if err != nil {
		log.Println(err)
		return
	}

	rate, err := GetCurrentNBURate()
	if err != nil {
		log.Println(err)
		return
	}

	for _, user := range users {
		err := sendEmail(user.Email, rate)
		if err != nil {
			log.Println(err)
		}
	}
}

func (env *Env) AddUser(w http.ResponseWriter, r *http.Request) {
	var user database.User

	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		http.Error(w, http.StatusText(400), 400)
		return
	}

	_, err = mail.ParseAddress(user.Email)
	if err != nil {
		http.Error(w, http.StatusText(400), 400)
		return
	}

	err = env.Users.AddUser(database.User{Email: user.Email})

	// check if "unique" constraint is violated
	var pqErr *pq.Error
	if errors.As(err, &pqErr) && pqErr.Code == "23505" {
		http.Error(w, http.StatusText(409), 409)
		return
	}
	if err != nil {
		http.Error(w, http.StatusText(400), 400)
		return
	}
}

func (env *Env) GetCurrentRate(w http.ResponseWriter, r *http.Request) {
	rate, err := GetCurrentNBURate()

	if err != nil {
		http.Error(w, http.StatusText(400), 400)
		return
	}

	w.Write([]byte(fmt.Sprintf("%f", rate)))
}
