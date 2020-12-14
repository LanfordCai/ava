package validator

import "github.com/filecoin-project/go-address"

// Filecoin ...
type Filecoin struct{}

var _ Validator = (*Filecoin)(nil)
var _ Prefixer = (*Filecoin)(nil)

// ValidateAddress returns validate result of arweave address
func (v *Filecoin) ValidateAddress(addr string, network NetworkType) *Result {
	if len(addr) < 3 || len(addr) > 86 {
		return &Result{false, Unknown, ""}
	}

	filAddr, err := address.NewFromString(addr)
	if err != nil {
		return &Result{false, Unknown, ""}
	}

	prefix := v.GetPrefix(network)
	if string(addr[0]) != prefix {
		return &Result{false, Unknown, ""}
	}

	addrType := FilSecp256k1
	switch filAddr.Protocol() {
	case address.ID:
		addrType = FilID
	case address.SECP256K1:
		addrType = FilSecp256k1
	case address.Actor:
		addrType = FilActor
	case address.BLS:
		addrType = FilBLS
	default:
		return &Result{false, Unknown, ""}
	}

	return &Result{true, addrType, ""}
}

// GetPrefix returns the prefix of filecoin address
func (v *Filecoin) GetPrefix(network NetworkType) string {
	if network == Mainnet {
		return "f"
	}
	return "t"
}
