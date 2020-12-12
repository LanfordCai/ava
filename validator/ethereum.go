package validator

import (
	"bytes"
	"encoding/hex"
	"strings"
	"unicode"
)

// Ethereum ..
type Ethereum struct{}

var _ Validator = (*Ethereum)(nil)
var _ Prefixer = (*Ethereum)(nil)

// ValidateAddress - Check an Ethereum address is valid or not
func (e *Ethereum) ValidateAddress(addr string, network NetworkType) *Result {
	if !e.isValidUncheckedAddress(addr) {
		return &Result{false, Unknown, ""}
	}

	if !e.withChecksum(addr) {
		return &Result{true, EthereumUnchecked, ""}
	}

	checksumedAddr := e.toChecksumedAddress(addr)
	if checksumedAddr != addr {
		return &Result{false, Unknown, ""}
	}

	return &Result{true, EthereumChecksumed, ""}
}

func (e *Ethereum) withChecksum(addr string) bool {
	noPrefixAddr := AddressWithoutPrefix(e, addr, Mainnet)

	return (strings.ToUpper(noPrefixAddr) != noPrefixAddr) &&
		(strings.ToLower(noPrefixAddr) != noPrefixAddr)
}

func (e *Ethereum) isValidUncheckedAddress(addr string) bool {
	if e.GetPrefix(Mainnet) != "0x" || len(addr) != 42 {
		return false
	}
	address := strings.ToLower(AddressWithoutPrefix(e, addr, Mainnet))
	if _, err := hex.DecodeString(address); err != nil {
		return false
	}
	return true
}

func (e *Ethereum) toChecksumedAddress(addr string) string {
	normalizedAddr := strings.ToLower(AddressWithoutPrefix(e, addr, Mainnet))
	hash := hex.EncodeToString(keccak256Sum([]byte(normalizedAddr)))
	var checksumAddrBuf bytes.Buffer
	for i, r := range normalizedAddr {
		if hash[i] >= '8' {
			checksumAddrBuf.Write([]byte{byte(unicode.ToUpper(r))})
		} else {
			checksumAddrBuf.Write([]byte{byte(r)})
		}
	}

	result := e.GetPrefix(Mainnet) + checksumAddrBuf.String()
	return result
}

// GetPrefix ...
func (e *Ethereum) GetPrefix(network NetworkType) string {
	return "0x"
}
