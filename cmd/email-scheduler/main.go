package main

import (
	"currency_exchange_rate/internal/api"
	"currency_exchange_rate/internal/database"
	"github.com/joho/godotenv"
	"log"
	"time"

	_ "github.com/lib/pq"
	"github.com/robfig/cron"
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

	utc, _ := time.LoadLocation("")
	cronHandler := cron.NewWithLocation(utc)
	cronHandler.AddFunc("* * * * * *", env.SendEmails)
	cronHandler.Start()
	select {}
}
