package validator

// BitcoinGold ...
type BitcoinGold struct{}

var _ BitcoinLike = (*BitcoinGold)(nil)

// ValidateAddress returns validate result of bitcoingold address
func (v *BitcoinGold) ValidateAddress(addr string, network NetworkType) *Result {
	if addrType := NormalAddrType(v, addr, network); addrType != Unknown {
		return &Result{IsValid: true, Type: addrType}
	}

	return &Result{IsValid: false, Type: Unknown}
}

// AddressVersion returns bitcoingold address version according to the address type and
// network type
func (v *BitcoinGold) AddressVersion(addrType AddressType, network NetworkType) byte {
	if network == Mainnet {
		if addrType == P2PKH {
			return 38
		}
		return 23
	}

	if addrType == P2PKH {
		return 111
	}
	return 196
}
