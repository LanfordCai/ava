package ava

import (
	"reflect"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestEthereumValidateAddress(t *testing.T) {
	validator := &Ethereum{}

	var validCases = map[string]*Result{
		"0x001d3f1ef827552ae1114027bd3ecf1f086ba0f9": {true, EthereumUnchecked, ""},
		"0xfb6916095ca1df60bb79ce92ce3ea74c37c5d359": {true, EthereumUnchecked, ""},
		"0x5aaeb6053f3e94c9b9a09f33669435e7ef1beaed": {true, EthereumUnchecked, ""},
		"0xdbf03b407c01e7cd3cbea99509d93f8dddc8c6fb": {true, EthereumUnchecked, ""},
		"0xd1220a0cf47c7b9be7a2e6ba89f429762e7b9adb": {true, EthereumUnchecked, ""},
		"0xFB6916095CA1DF60BB79CE92CE3EA74C37C5D359": {true, EthereumUnchecked, ""},
		"0x001D3F1EF827552AE1114027BD3ECF1F086BA0F9": {true, EthereumUnchecked, ""},
		"0x5AAEB6053F3E94C9B9A09F33669435E7EF1BEAED": {true, EthereumUnchecked, ""},
		"0xDBF03B407C01E7CD3CBEA99509D93F8DDDC8C6FB": {true, EthereumUnchecked, ""},
		"0xD1220A0CF47C7B9BE7A2E6BA89F429762E7B9ADB": {true, EthereumUnchecked, ""},
		"0x0F53Ec6bbd2B6712c07d8880E5C8f08753d0d5D5": {true, EthereumChecksumed, ""},
		"0x001d3F1ef827552Ae1114027BD3ECF1f086bA0F9": {true, EthereumChecksumed, ""},
		"0xfB6916095ca1df60bB79Ce92cE3Ea74c37c5d359": {true, EthereumChecksumed, ""},
		"0x5aAeb6053F3E94C9b9A09f33669435E7Ef1BeAed": {true, EthereumChecksumed, ""},
		"0xdbF03B407c01E7cD3CBea99509d93f8DDDC8C6FB": {true, EthereumChecksumed, ""},
		"0xD1220A0cf47c7B9Be7A2E6BA89F429762e7b9aDb": {true, EthereumChecksumed, ""},
	}

	for addr, result := range validCases {
		assert.True(t, reflect.DeepEqual(validator.ValidateAddress(addr, Mainnet), result), addr)
		assert.True(t, reflect.DeepEqual(validator.ValidateAddress(addr, Testnet), result), addr)
	}

	var invalidCases = map[string]*Result{
		"0x001d3F1ef827552Ae1114027BD3ECF1F086bA0F9": {false, Unknown, ""},
		"0xfB6916095ca1df60bB79Ce92cE3Ea74C37c5d359": {false, Unknown, ""},
		"0x5aAeb6053F3E94C9b9A09f33669435E7ef1BeAed": {false, Unknown, ""},
		"0xd1220A0cf47c7B9Be7A2E6BA89F429762e7b9aDb": {false, Unknown, ""},
		"0xdbF03B407c01E7cD3CBea99509d93f8DDDC8C6Fb": {false, Unknown, ""},
		"": {false, Unknown, ""},
	}

	for addr, result := range invalidCases {
		assert.True(t, reflect.DeepEqual(validator.ValidateAddress(addr, Testnet), result), addr)
		assert.True(t, reflect.DeepEqual(validator.ValidateAddress(addr, Mainnet), result), addr)
	}
}
