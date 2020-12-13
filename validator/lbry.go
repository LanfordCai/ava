package validator

import "github.com/btcsuite/btcutil/base58"

// Lbry ...
type Lbry struct{}

var _ BitcoinLike = (*Lbry)(nil)

// ValidateAddress returns validate result of qtum address
// NOTICE: only support mainnet P2PKH now
func (v *Lbry) ValidateAddress(addr string, network NetworkType) *Result {
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
func (v *Lbry) AddressVersion(addrType AddressType, network NetworkType) byte {
	if addrType == P2PKH && network == Mainnet {
		return 85
	}
	panic(ErrUnsupported.Error())
}
