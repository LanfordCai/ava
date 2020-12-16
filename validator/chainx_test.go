package validator

import (
	"reflect"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestChainXValidateAddress(t *testing.T) {
	validator := &ChainX{}

	var validCases = map[string]*Result{
		"5QxFTz6oeBbgPtGMg531zUZ4JJz2RkbiCjkUu9z1Uf7Wi92d": {Success, true, Normal, ""},
		"5SdW3TbRpnsbvTDa9QDXftM8mtdw86ou8Z1Lyd3D4XgjSEii": {Success, true, Normal, ""},
		"5QbeyxMUMhEeukBaxAqHip9E1u93kMqxf2dTBPDrPeY52B9o": {Success, true, Normal, ""},
		"5R7sxzEYGGSGSCQd3DQLyYwkA29jYGYaecHUHCsQ8s1Vxeo7": {Success, true, Normal, ""},
		"5VQLcbTNn6UPYqLK87kPBHQvhvYL24z9bCa9XBtkkWekGrwj": {Success, true, Normal, ""},
		"5UBLKqQ7efPHxwVQCT82Vf86PxsDRCRv8nVZZQ9YpUrkg25M": {Success, true, Normal, ""},
	}

	for addr, result := range validCases {
		assert.True(t, reflect.DeepEqual(validator.ValidateAddress(addr, Mainnet), result), addr)
	}

	var invalidCases = map[string]*Result{
		"15SdW3TbRpnsbvTDa9QDXftM8mtdw86ou8Z1Ly3D4XgjSEii":  {Success, false, Unknown, ""},
		"15j4dg5GzsL1bw2U2AWgeyAk6QTxq43V7ZPbXdAmbVLjvDCK":  {Success, false, Unknown, ""},
		"6VQLcbTNn6UPYqLK87kPBHQvhvYL24z9bCa9XBtkkWekGrwj":  {Success, false, Unknown, ""},
		"5VQLcbTNn6UPYqLK87kPBHQvhvYL24z9bCa9XBtkkWekGewj":  {Success, false, Unknown, ""},
		"5VQLcbTNn6UPYqLK87kPBHQvhvML24z9bCa9XBtkkWekGrwj ": {Success, false, Unknown, ""},
		"bc1q3l9k4lm5z4mtsl6smmj9qxy03e65x3maz4p9xv":        {Success, false, Unknown, ""},
		"abcde": {Success, false, Unknown, ""},
		"":      {Success, false, Unknown, ""},
	}

	for addr, result := range invalidCases {
		assert.True(t, reflect.DeepEqual(validator.ValidateAddress(addr, Mainnet), result), addr)
		assert.True(t, reflect.DeepEqual(validator.ValidateAddress(addr, Testnet), result), addr)
	}
}
