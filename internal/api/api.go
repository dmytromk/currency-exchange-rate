package api

import (
	"currency_exchange_rate/internal/database"
	"encoding/json"
	"fmt"
	"net/http"
	"net/mail"
)

type Env struct {
	Users database.UserModel
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
