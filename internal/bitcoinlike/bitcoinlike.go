package bitcoinlike

import (
	"bytes"
	"fmt"

	"github.com/LanfordCai/ava/internal/common"
	"github.com/LanfordCai/ava/internal/utils"
	"github.com/btcsuite/btcutil/base58"
	"github.com/btcsuite/btcutil/bech32"
	"github.com/pkg/errors"
)

const (
	programLenP2WPKH = 20
	programLenP2WSH  = 32
)

// Validator - Method receiver, used to validate bitcoinlike addresses
type Validator struct {
	ChainName           string
	MainnetBech32HRP    string
	TestnetBech32HRP    string
	MainnetP2PKHAddrVer []byte
	MainnetP2SHAddrVer  []byte
	TestnetP2PKHAddrVer []byte
	TestnetP2SHAddrVer  []byte
	EnabledTypes        []string
	SupportedTypes      []string
}

// ValidateAddress - Validate the given address and network type
// SEE: https://en.bitcoin.it/wiki/List_of_address_prefixes
func (b *Validator) ValidateAddress(address string, isTestnet bool) common.ValidationResult {
	var addrType string
	switch {
	case b.isP2PKH(address, isTestnet):
		addrType = "P2PKH"
	case b.isP2SH(address, isTestnet):
		addrType = "P2SH"
	case b.isP2WPKH(address, isTestnet):
		addrType = "P2WPKH"
	case b.isP2WSH(address, isTestnet):
		addrType = "P2WSH"
	default:
		addrType = "Unknown"
	}

	valid := false
	if addrType != "Unknown" && utils.Contains(b.EnabledTypes, addrType) {
		valid = true
	}
	return common.NewValidationResult(valid, addrType)
}

// CheckTypes - Check address type has been supported or not
func (b *Validator) CheckTypes(types []string) error {
	if len(types) < 1 {
		return common.ErrEmptyEnabledTypes
	}

	for _, t := range types {
		if !utils.Contains(b.SupportedTypes, t) {
			errMsg := fmt.Sprintf("Supported types are: %+q", b.SupportedTypes)
			return errors.Wrap(common.ErrUnsupportedTypes, errMsg)
		}
	}

	return nil
}

func (b *Validator) isP2PKH(address string, isTestnet bool) bool {
	decoded, version, err := base58.CheckDecode(address)
	if err != nil {
		return false
	}

	if b.ChainName == "Zcash" {
		if isTestnet {
			return bytes.Compare([]byte{version, decoded[0]}, b.TestnetP2PKHAddrVer) == 0
		}
		return bytes.Compare([]byte{version, decoded[0]}, b.MainnetP2PKHAddrVer) == 0
	}

	if isTestnet {
		return bytes.Compare([]byte{version}, b.TestnetP2PKHAddrVer) == 0
	}
	return bytes.Compare([]byte{version}, b.MainnetP2PKHAddrVer) == 0
}

func (b *Validator) isP2SH(address string, isTestnet bool) bool {
	decoded, version, err := base58.CheckDecode(address)
	if err != nil {
		return false
	}

	if b.ChainName == "Zcash" {
		if isTestnet {
			return bytes.Compare([]byte{version, decoded[0]}, b.TestnetP2SHAddrVer) == 0
		}
		return bytes.Compare([]byte{version, decoded[0]}, b.MainnetP2SHAddrVer) == 0
	}

	if isTestnet {
		return bytes.Compare([]byte{version}, b.TestnetP2SHAddrVer) == 0
	}
	return bytes.Compare([]byte{version}, b.MainnetP2SHAddrVer) == 0
}

func (b *Validator) isP2WPKH(address string, isTestnet bool) bool {
	hrp, version, program, err := segwitAddrDecode(address)
	if err != nil || version != 0 || len(program) != programLenP2WPKH {
		return false
	}
	return (isTestnet && hrp == b.TestnetBech32HRP) || (!isTestnet && hrp == b.MainnetBech32HRP)
}

func (b *Validator) isP2WSH(address string, isTestnet bool) bool {
	hrp, version, program, err := segwitAddrDecode(address)
	if err != nil || version != 0 || len(program) != programLenP2WSH {
		return false
	}
	return (isTestnet && hrp == b.TestnetBech32HRP) || (!isTestnet && hrp == b.MainnetBech32HRP)
}

func segwitAddrDecode(address string) (string, byte, []byte, error) {
	hrp, data, err := bech32.Decode(address)
	if err != nil || len(data) < 1 {
		return "", 255, []byte{}, err
	}
	version := data[0]
	program, err := bech32.ConvertBits(data[1:], 5, 8, false)
	if err != nil {
		return "", 255, []byte{}, err
	}
	return hrp, version, program, nil
}
