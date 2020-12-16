package main

import (
	"fmt"
	"os"

	"ava/validator"
)

func validateBitcoinAddr() {
	v := &validator.Bitcoin{}
	addr := "19JeUHUvw23fwKeK1zZD4moKyxj1xn4Kxi"
	result := v.ValidateAddress(addr, validator.Mainnet)

	fmt.Printf("Address is valid?: %t\n", result.IsValid)
	fmt.Printf("Address type: %s\n", result.Type)
}

func validateEOSAddr() {
	client := validator.EOSClient{
		Endpoint: os.Getenv("AVA_EOS_ENDPOINT"),
	}
	v := &validator.EOS{Client: &client}

	result := v.ValidateAddress("huobideposit", validator.Mainnet)
	if result.Status == validator.Success {
		fmt.Printf("Address is valid?: %t\n", result.IsValid)
		fmt.Printf("Address type: %s\n", result.Type)
	} else {
		fmt.Printf("Address validation failed: %s\n", result.Msg)
	}

	deployed, err := v.IsContractDeployed("huobideposit")
	if err != nil {
		fmt.Printf("Get deployed info failed: %s\n", err.Error())
	} else {
		fmt.Printf("Contract is deployed?: %t\n", deployed)
	}
}

func main() {
	fmt.Println("Validate Bitcoin Address:")
	validateBitcoinAddr()

	fmt.Println("\nValidate EOS Address:")
	validateEOSAddr()
}
