package validator

// Dash ...
type Dash struct{}

var _ BitcoinLike = (*Dash)(nil)

// ValidateAddress returns validate result of dash address
func (v *Dash) ValidateAddress(addr string, network NetworkType) *Result {
	if addrType := NormalAddrType(v, addr, network); addrType != Unknown {
		return &Result{Success, true, addrType, ""}
	}

	return &Result{Success, false, Unknown, ""}
}

// AddressVersion returns dash address version according to the address type and
// network type
func (v *Dash) AddressVersion(addrType AddressType, network NetworkType) byte {
	if network == Mainnet {
		if addrType == P2PKH {
			return 76
		}
		return 16
	}

	if addrType == P2PKH {
		return 140
	}
	return 19
}
