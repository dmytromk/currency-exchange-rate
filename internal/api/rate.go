package api

import (
	"encoding/json"
	"net/http"
	"time"
)

type ResponseNBU struct {
	R030         int     `json:"r030"`
	TXT          string  `json:"txt"`
	Rate         float32 `json:"rate"`
	CC           string  `json:"cc"`
	ExchangeDate string  `json:"exchangedate"`
}

func GetCurrentNBURate() (float32, error) {
	var myClient = &http.Client{Timeout: 10 * time.Second}
	var result []ResponseNBU

	r, err := myClient.Get("https://bank.gov.ua/NBUStatService/v1/statdirectory/dollar_info?json")
	if err != nil {
		return -1, err
	}

	err = json.NewDecoder(r.Body).Decode(&result)
	if err != nil {
		return -1, err
	}
	return result[0].Rate, nil
}
