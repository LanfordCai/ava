package ava

import (
	"github.com/btcsuite/btcutil/base58"
)

// Neo ...
type Neo struct{}

var _ ValidatorBitcoinLike = (*Neo)(nil)

// ValidateAddress returns validate result of neo address
func (v *Neo) ValidateAddress(addr string, network NetworkType) *Result {
	decoded, version, err := base58.CheckDecode(addr)
	if err != nil || len(decoded) != 20 {
		return &Result{false, Unknown, ""}
	}

	expectP2PKH := v.AddressVersion(P2PKH, network)
	if version == expectP2PKH {
		return &Result{true, P2PKH, ""}
	}

	return &Result{false, Unknown, ""}
}

// AddressVersion returns neo address version according to the address type and
// network type
func (v *Neo) AddressVersion(addrType AddressType, network NetworkType) byte {
	if network == Mainnet && addrType == P2PKH {
		return 23
	}
	panic(ErrUnsupported.Error())
}

// AddressHrp returns hrps of neo according to the network
func (v *Neo) AddressHrp(network NetworkType) string {
	panic(ErrUnsupported.Error())
}

// SegwitProgramLength returns segwit program length of neo
func (v *Neo) SegwitProgramLength(addrType AddressType) int {
	panic(ErrUnsupported.Error())
}
