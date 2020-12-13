package validator

import "ava/cashaddr"

// BitcoinCash ...
type BitcoinCash struct{}

var _ BitcoinLike = (*BitcoinCash)(nil)
var _ CashAddress = (*BitcoinCash)(nil)

// ValidateAddress returns validate result of bitcoincash address
func (v *BitcoinCash) ValidateAddress(addr string, network NetworkType) *Result {
	if addrType := NormalAddrType(v, addr, network); addrType != Unknown {
		return &Result{true, addrType, ""}
	}

	if addrType := v.CashAddrType(addr, network); addrType != Unknown {
		return &Result{true, addrType, ""}
	}

	return &Result{false, Unknown, ""}
}

// AddressVersion returns bitcoincash address version according to the address type and
// network type
func (v *BitcoinCash) AddressVersion(addrType AddressType, network NetworkType) byte {
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

// CashAddrType ...
func (v *BitcoinCash) CashAddrType(addr string, network NetworkType) AddressType {
	legacyAddr, err := cashaddr.ToLegacyAddr(addr)
	// if not cashaddr
	if err != nil {
		return Unknown
	}

	addrType := NormalAddrType(v, legacyAddr, network)

	if addrType == P2PKH {
		return CashAddrP2PKH
	} else if addrType == P2SH {
		return CashAddrP2SH
	}
	return Unknown
}
