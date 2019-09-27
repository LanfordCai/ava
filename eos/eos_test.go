package eos

// import (
// 	"testing"

// 	"github.com/stretchr/testify/assert"
// )

// var validAddresses = []string{
// 	"binancecold1",
// 	"bithumbshiny",
// 	"bitfinexcw55",
// 	"wnnqcyubjoeo",
// 	"big.one",
// }

// var invalidAddresses = []string{
// 	"0inancecold1",
// 	"abc-decwdsae",
// 	"binancecold12",
// 	" ",
// 	"",
// }

// func TestIsValidAddress(t *testing.T) {
// 	validator := New(false)
// 	for _, a := range validAddresses {
// 		isValid := validator.IsValidAddress(a, false)
// 		assert.Equal(t, true, isValid)
// 	}

// 	for _, a := range invalidAddresses {
// 		isValid := validator.IsValidAddress(a, false)
// 		assert.Equal(t, false, isValid)
// 	}
// }
