package bitcoinlike

import (
	"github.com/LanfordCai/ava/utils"
	"github.com/btcsuite/btcutil/base58"
)

// Validator - Method receiver, used to validate bitcoinlike addresses
type Validator struct {
	MainnetP2PKHAddrVer byte
	MainnetP2SHAddrVer  byte
	TestnetP2PKHAddrVer byte
	TestnetP2SHAddrVer  byte
	AcceptableTypes     []string
}

// IsValidAddress - Validate the given address and network type
// SEE: https://en.bitcoin.it/wiki/List_of_address_prefixes
func (b *Validator) IsValidAddress(address string, isTestnet bool) bool {
	if b.isP2PKH(address, isTestnet) && utils.Contains(b.AcceptableTypes, "P2PKH") {
		return true
	}

	if b.isP2SH(address, isTestnet) && utils.Contains(b.AcceptableTypes, "P2SH") {
		return true
	}

	return false
}

func (b *Validator) isP2PKH(address string, isTestnet bool) bool {
	_, version, err := base58.CheckDecode(address)
	if err != nil {
		return false
	}
	if isTestnet {
		return version == b.TestnetP2PKHAddrVer
	}
	return version == b.MainnetP2PKHAddrVer
}

func (b *Validator) isP2SH(address string, isTestnet bool) bool {
	_, version, err := base58.CheckDecode(address)
	if err != nil {
		return false
	}
	if isTestnet {
		return version == b.TestnetP2SHAddrVer
	}
	return version == b.MainnetP2SHAddrVer
}
