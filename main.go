package main

import (
	"currency_exchange_rate/internal/api"
	"fmt"
)

func main() {
	fmt.Println(api.GetCurrentNBURate())
}
