package main

import (
	"fmt"

	"frontendmasters.com/go/crypto/api"
)

func main() {
	rate, err := api.GetRate("BTC")

	if err == nil {
		fmt.Printf("The rate of %v is %.2f\n", rate.Currency, rate.Price)
	}
	print(rate.Price, err)
}
