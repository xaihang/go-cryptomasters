package main

import (
	"fmt"
	"sync"

	"frontendmasters.com/go/crypto/api"
)

func main() {
	currencies := []string{"BTC", "ETH", "BCH"}
	var wg sync.WaitGroup
	for _, currency := range currencies {
		wg.Add(1)
		go func(currencyCode string) {
			getCurrencyData(currencyCode)
			wg.Done()
		}(currency)
	}
}

func getCurrencyData(currency string) {
	rate, err := api.GetRate(currency)
	if err == nil {
		fmt.Printf("The rate of %v is %.2f\n", rate.Currency, rate.Price)
	}
}
