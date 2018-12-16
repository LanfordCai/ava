package bytom

import (
	"github.com/btcsuite/btcutil/bech32"
)

// IsValidAddress ...
func IsValidAddress(address string, isTestnet bool) bool {
	hrp, version, program, err := segwitAddrDecode(address)
	if err != nil || version != 0 || (len(program) != 32 && len(program) != 20) {
		return false
	}
	return (isTestnet && hrp == "tm") || (!isTestnet && hrp == "bm")
}

// IsP2WPKH ...
func IsP2WPKH(address string, isTestnet bool) bool {
	hrp, version, program, err := segwitAddrDecode(address)
	if err != nil || version != 0 || len(program) != 20 {
		return false
	}
	return (isTestnet && hrp == "tm") || (!isTestnet && hrp == "bm")
}

// IsP2WSH ...
func IsP2WSH(address string, isTestnet bool) bool {
	hrp, version, program, err := segwitAddrDecode(address)
	if err != nil || version != 0 || len(program) != 32 {
		return false
	}
	return (isTestnet && hrp == "tm") || (!isTestnet && hrp == "bm")
}

func segwitAddrDecode(address string) (string, byte, []byte, error) {
	hrp, data, err := bech32.Decode(address)
	if err != nil || len(data) < 1 {
		return "", 255, []byte{}, err
	}
	version := data[0]
	program, err := bech32.ConvertBits(data[1:], 5, 8, false)
	if err != nil {
		return "", 255, []byte{}, err
	}
	return hrp, version, program, nil
}
