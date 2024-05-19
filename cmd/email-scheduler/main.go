package main

import (
	"fmt"
	"time"

	_ "github.com/lib/pq"
	"github.com/robfig/cron"
)

func main() {
	utc, _ := time.LoadLocation("")
	cronHandler := cron.NewWithLocation(utc)
	cronHandler.AddFunc("* * * * * *", func() {
		fmt.Println("crontab ping")
	})
	cronHandler.Start()
	select {}
}
