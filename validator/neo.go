package validator

import (
	"github.com/btcsuite/btcutil/base58"
)

// NEO ...
type NEO struct{}

var _ BitcoinLike = (*NEO)(nil)

// ValidateAddress returns validate result of NEO address
func (v *NEO) ValidateAddress(addr string, network NetworkType) *Result {
	decoded, version, err := base58.CheckDecode(addr)
	if err != nil || len(decoded) != 20 {
		return &Result{Success, false, Unknown, ""}
	}

	expectP2PKH := v.AddressVersion(P2PKH, network)
	if version == expectP2PKH {
		return &Result{Success, true, P2PKH, ""}
	}

	return &Result{Success, false, Unknown, ""}
}

// AddressVersion returns NEO address version according to the address type and
// network type
func (v *NEO) AddressVersion(addrType AddressType, network NetworkType) byte {
	if network == Mainnet && addrType == P2PKH {
		return 23
	}
	panic(ErrUnsupported.Error())
}
