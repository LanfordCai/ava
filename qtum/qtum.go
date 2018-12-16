package qtum

import (
	"github.com/btcsuite/btcutil/base58"
)

// IsValidAddress ...
// SEE: https://github.com/qtumproject/qtumcore-lib/blob/master/lib/networks.js
func IsValidAddress(address string, isTestnet bool) bool {
	return IsP2PKH(address, isTestnet) || IsP2SH(address, isTestnet)
}

// IsP2PKH ...
func IsP2PKH(address string, isTestnet bool) bool {
	_, version, err := base58.CheckDecode(address)
	if err != nil {
		return false
	}
	if isTestnet {
		return version == 120
	}
	return version == 58
}

// IsP2SH ...
func IsP2SH(address string, isTestnet bool) bool {
	_, version, err := base58.CheckDecode(address)
	if err != nil {
		return false
	}
	if isTestnet {
		return version == 110
	}
	return version == 50
}
