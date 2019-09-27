package eos

// import (
// 	"regexp"
// )

// // Validator - EOS address validator
// type Validator struct {
// 	allowContract bool
// }

// // New - Create a new EOS address validator
// func New(allowContract bool) *Validator {
// 	return &Validator{allowContract: allowContract}
// }

// // ValidateAddress ...
// func (e *Validator) ValidateAddress(address string, isTestnet bool) bool {
// 	if withRightAddressFormat(address) && isRegistered(address) {
// 		if !e.allowContract && isContract(address) {
// 			return false
// 		}

// 		return true
// 	}

// 	return false
// }

// func isRegistered(address string) bool {
// 	return true
// }

// func isContract(address string) bool {
// 	return true
// }

// func withRightAddressFormat(address string) bool {
// 	matched, err := regexp.MatchString("[^.12345abcdefghijklmnopqrstuvwxyz]", address)
// 	if err != nil || len(address) > 12 || len(address) == 0 {
// 		return false
// 	}
// 	return !matched
// }
