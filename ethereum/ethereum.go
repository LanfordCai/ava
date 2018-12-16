// https://ethereum.stackexchange.com/questions/1374/how-can-i-check-if-an-ethereum-address-is-valid
package ethereum

import (
	"bytes"
	"encoding/hex"
	"errors"
	"strings"
	"unicode"

	"github.com/ethereum/go-ethereum/crypto/sha3"
)

// IsValidAddress ...
func IsValidAddress(address string) bool {
	if isValidNonChecksumAddress(address) {
		noPrefixAddr := address[2:]
		if (strings.ToUpper(noPrefixAddr) == noPrefixAddr) || (strings.ToLower(noPrefixAddr) == noPrefixAddr) {
			return true
		}
		checksumAddress, err := ToChecksumAddress(address)
		if err != nil || checksumAddress != address {
			return false
		}
		return true
	}
	return false
}

// ToChecksumAddress ...
func ToChecksumAddress(address string) (string, error) {
	if !isValidNonChecksumAddress(address) {
		return "", errors.New("Invalid address")
	}

	address = strings.ToLower(address[2:])
	hash := keccak256Sum(address)
	var checksumAddrBuf bytes.Buffer
	for i, r := range address {
		if hash[i] >= '8' {
			checksumAddrBuf.Write([]byte{byte(unicode.ToUpper(r))})
		} else {
			checksumAddrBuf.Write([]byte{byte(r)})
		}
	}

	result := checksumAddrBuf.String()
	return "0x" + result, nil
}

func keccak256Sum(address string) string {
	hash := sha3.NewKeccak256()
	var buf []byte
	hash.Write([]byte(address))
	buf = hash.Sum(buf)

	return hex.EncodeToString(buf)
}

func isValidNonChecksumAddress(address string) bool {
	if !strings.HasPrefix(address, "0x") || len(address) != 42 {
		return false
	}
	address = strings.ToLower(address[2:])
	if _, err := hex.DecodeString(address); err != nil {
		return false
	}
	return true
}
