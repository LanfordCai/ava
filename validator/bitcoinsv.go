package validator

// BitcoinSV ...
type BitcoinSV struct{}

var _ BitcoinLike = (*BitcoinSV)(nil)

// ValidateAddress returns validate result of bitcoin address
func (v *BitcoinSV) ValidateAddress(addr string, network NetworkType) *Result {
	if addrType := NormalAddrType(v, addr, network); addrType != Unknown {
		return &Result{Success, true, addrType, ""}
	}

	return &Result{Success, false, Unknown, ""}
}

// AddressVersion returns bitcoin address version according to the address type and
// network type
func (v *BitcoinSV) AddressVersion(addrType AddressType, network NetworkType) byte {
	if network == Mainnet {
		if addrType == P2PKH {
			return 0
		}
		return 5
	}

	if addrType == P2PKH {
		return 111
	}
	return 196
}
