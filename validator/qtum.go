package validator

// Qtum ...
type Qtum struct{}

var _ BitcoinLike = (*Qtum)(nil)

// ValidateAddress returns validate result of qtum address
func (v *Qtum) ValidateAddress(addr string, network NetworkType) *Result {
	if addrType := NormalAddrType(v, addr, network); addrType != Unknown {
		return &Result{Success, true, addrType, ""}
	}

	return &Result{Success, false, Unknown, ""}
}

// AddressVersion returns qtum address version according to the address type and
// network type
func (v *Qtum) AddressVersion(addrType AddressType, network NetworkType) byte {
	if network == Mainnet {
		if addrType == P2PKH {
			return 58
		}
		return 50
	}

	if addrType == P2PKH {
		return 120
	}
	return 110
}
