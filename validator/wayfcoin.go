package validator

import "github.com/btcsuite/btcutil/base58"

// Wayfcoin ...
type Wayfcoin struct{}

var _ BitcoinLike = (*Wayfcoin)(nil)

// ValidateAddress returns validate result of qtum address
// NOTICE: only support mainnet P2PKH now
func (v *Wayfcoin) ValidateAddress(addr string, network NetworkType) *Result {
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

// AddressVersion returns qtum address version according to the address type and
// network type
func (v *Wayfcoin) AddressVersion(addrType AddressType, network NetworkType) byte {
	if addrType == P2PKH && network == Mainnet {
		return 73
	}
	panic(ErrUnsupported.Error())
}
