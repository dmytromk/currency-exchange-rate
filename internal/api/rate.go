package api

import (
	"encoding/json"
	"io"
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

func getCurrentNBURate() (float32, error) {
	var myClient = &http.Client{Timeout: 10 * time.Second}

	response, err := myClient.Get("https://bank.gov.ua/NBUStatService/v1/statdirectory/dollar_info?json")

	if err != nil {
		return -1, err
	}

	responseData, err := io.ReadAll(response.Body)
	if err != nil {
		return -1, err
	}

	var result []ResponseNBU
	err = json.Unmarshal(responseData, &result)
	if err != nil {
		return -1, err
	}
	return result[0].Rate, nil
}
