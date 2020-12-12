package main

import (
	"fmt"

	"github.com/LanfordCai/ava"
)

func main() {
	validator := &ava.Bitcoin{}
	addr := "19JeUHUvw23fwKeK1zZD4moKyxj1xn4Kxi"
	result := validator.ValidateAddress(addr, ava.Mainnet)

	fmt.Printf("Address is valid?: %t\n", result.IsValid)
	fmt.Printf("Address type: %s\n", result.Type)
}
