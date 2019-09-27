package ethereum

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestValidateAddress(t *testing.T) {
	validator := New()
	var validAddresses = []string{
		"0x001d3f1ef827552ae1114027bd3ecf1f086ba0f9",
		"0xfb6916095ca1df60bb79ce92ce3ea74c37c5d359",
		"0x5aaeb6053f3e94c9b9a09f33669435e7ef1beaed",
		"0xdbf03b407c01e7cd3cbea99509d93f8dddc8c6fb",
		"0xd1220a0cf47c7b9be7a2e6ba89f429762e7b9adb",
		"0x001d3F1ef827552Ae1114027BD3ECF1f086bA0F9",
		"0xfB6916095ca1df60bB79Ce92cE3Ea74c37c5d359",
		"0x5aAeb6053F3E94C9b9A09f33669435E7Ef1BeAed",
		"0xdbF03B407c01E7cD3CBea99509d93f8DDDC8C6FB",
		"0xD1220A0cf47c7B9Be7A2E6BA89F429762e7b9aDb",
		"0x001D3F1EF827552AE1114027BD3ECF1F086BA0F9",
		"0xFB6916095CA1DF60BB79CE92CE3EA74C37C5D359",
		"0x5AAEB6053F3E94C9B9A09F33669435E7EF1BEAED",
		"0xDBF03B407C01E7CD3CBEA99509D93F8DDDC8C6FB",
		"0xD1220A0CF47C7B9BE7A2E6BA89F429762E7B9ADB",
		"0x0F53Ec6bbd2B6712c07d8880E5C8f08753d0d5D5",
	}

	var invalidAddresses = []string{
		"0x001d3F1ef827552Ae1114027BD3ECF1F086bA0F9",
		"0xfB6916095ca1df60bB79Ce92cE3Ea74C37c5d359",
		"0x5aAeb6053F3E94C9b9A09f33669435E7ef1BeAed",
		"0xdbF03B407c01E7cD3CBea99509d93f8DDDC8C6Fb",
		"0xd1220A0cf47c7B9Be7A2E6BA89F429762e7b9aDb",
		"",
	}

	for _, a := range validAddresses {
		result := validator.ValidateAddress(a, false)
		assert.Equal(t, true, result.IsValid)
	}

	for _, a := range invalidAddresses {
		result := validator.ValidateAddress(a, false)
		assert.Equal(t, false, result.IsValid)
	}
}

func TestToChecksumAddress(t *testing.T) {
	var checkSumAddressPairs = []struct {
		address         string
		checksumAddress string
	}{
		{"0x001d3f1ef827552ae1114027bd3ecf1f086ba0f9", "0x001d3F1ef827552Ae1114027BD3ECF1f086bA0F9"},
		{"0xfb6916095ca1df60bb79ce92ce3ea74c37c5d359", "0xfB6916095ca1df60bB79Ce92cE3Ea74c37c5d359"},
		{"0x5aaeb6053f3e94c9b9a09f33669435e7ef1beaed", "0x5aAeb6053F3E94C9b9A09f33669435E7Ef1BeAed"},
		{"0xdbf03b407c01e7cd3cbea99509d93f8dddc8c6fb", "0xdbF03B407c01E7cD3CBea99509d93f8DDDC8C6FB"},
		{"0xd1220a0cf47c7b9be7a2e6ba89f429762e7b9adb", "0xD1220A0cf47c7B9Be7A2E6BA89F429762e7b9aDb"},
	}

	validator := New()

	for _, p := range checkSumAddressPairs {
		checksumed, err := validator.ToChecksumAddress(p.address)
		if err != nil {
			t.Logf("Failed for %s", p.address)
		} else {
			assert.Equal(t, true, checksumed == p.checksumAddress)
		}
	}
}
