package api

import (
	"currency_exchange_rate/internal/database"
	"encoding/binary"
	"encoding/json"
	"fmt"
	"github.com/lib/pq"
	"math"
	"net/http"
	"net/mail"
)

type Env struct {
	Users database.UserModel
}

func (env *Env) SendEmails() {
	users, err := env.Users.GetAllUsers()
	if err != nil {
		fmt.Println(err)
		return
	}

	rate, err := GetCurrentNBURate()
	if err != nil {
		fmt.Println(err)
		return
	}

	for _, user := range users {
		to := []string{user.Email}

		var msg []byte
		binary.LittleEndian.PutUint32(msg[:], math.Float32bits(rate))

		err := sendEmail(to, msg)
		if err != nil {
			fmt.Println(err)
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
	pqErr := err.(*pq.Error)
	// check if "unique" constraint is violated
	if pqErr.Code == "23505" {
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
