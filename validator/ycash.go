package validator

// Ycash ...
type Ycash struct{}

var _ ZcashLike = (*Ycash)(nil)

// ValidateAddress returns validate result of ycash address
func (v *Ycash) ValidateAddress(addr string, network NetworkType) *Result {
	if addrType := ZcashlikeNormalAddrType(v, addr, network); addrType != Unknown {
		return &Result{IsValid: true, Type: addrType}
	}

	return &Result{IsValid: false, Type: Unknown}
}

// AddressVersion returns ycash address version according to the address type and
// network type
func (v *Ycash) AddressVersion(addrType AddressType, network NetworkType) []byte {
	if network == Mainnet {
		if addrType == P2PKH {
			return []byte{28, 40}
		}
		return []byte{28, 44}
	}

	if addrType == P2PKH {
		return []byte{28, 149}
	}
	return []byte{28, 42}
}
