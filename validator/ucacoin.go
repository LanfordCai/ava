package validator

import "github.com/btcsuite/btcutil/base58"

// Ucacoin ...
type Ucacoin struct{}

var _ BitcoinLike = (*Ucacoin)(nil)

// ValidateAddress returns validate result of qtum address
// NOTICE: only support mainnet P2PKH now
func (v *Ucacoin) ValidateAddress(addr string, network NetworkType) *Result {
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

// AddressVersion returns qtum address version according to the address type and
// network type
func (v *Ucacoin) AddressVersion(addrType AddressType, network NetworkType) byte {
	if addrType == P2PKH && network == Mainnet {
		return 68
	}
	panic(ErrUnsupported.Error())
}
