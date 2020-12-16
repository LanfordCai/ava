package validator

import (
	"github.com/LanfordCai/ava/base58check"
)

// Vsystems ...
type Vsystems struct{}

var _ Validator = (*Vsystems)(nil)

const vsystemsAddrVersion = 5

// ValidateAddress returns validate result of vsystems address
func (v *Vsystems) ValidateAddress(addr string, network NetworkType) *Result {
	encoder := base58check.BitcoinEncoder
	encoder.ChecksumType = base58check.ChecksumBlake2bKeccak256

	decoded, err := encoder.CheckDecode(addr)
	if err != nil {
		return &Result{Success, false, Unknown, ""}
	}

	version := decoded[0]
	networkByte := decoded[1]

	if version == vsystemsAddrVersion && networkByte == v.getNetworkByte(network) {
		return &Result{Success, true, Normal, ""}
	}

	return &Result{Success, false, Unknown, ""}
}

func (v *Vsystems) getNetworkByte(network NetworkType) byte {
	if network == Mainnet {
		return 77
	}
	return 84
}
