package validator

import (
	"bytes"
	"encoding/base32"
	"encoding/binary"

	"github.com/snksoft/crc"
)

// Stellar ...
type Stellar struct {
	Client *StellarClient
}

var _ OnchainValidator = (*Stellar)(nil)

const stellarPrefix = 48
const stellarChecksumLen = 2
const stellarAddrLen = 56

// ValidateAddress returns validate result of stellar address
func (v *Stellar) ValidateAddress(addr string, network NetworkType) *Result {
	if isValid := v.IsAddrFormatValid(addr); !isValid {
		return &Result{Success, false, Unknown, ""}
	}

	addrType, err := v.Client.GetAccount(addr)
	if err != nil {
		return &Result{Failure, false, Unknown, err.Error()}
	}

	if addrType == Unknown {
		return &Result{Success, false, Unknown, ""}
	}

	return &Result{Success, true, addrType, ""}
}

// IsAddrFormatValid ...
func (v *Stellar) IsAddrFormatValid(addr string) bool {
	if len(addr) != stellarAddrLen {
		return false
	}

	data, err := base32.StdEncoding.DecodeString(addr)
	if err != nil {
		return false
	}
	payload := data[:len(data)-stellarChecksumLen]
	if payload[0] != stellarPrefix {
		return false
	}

	checksum := data[len(data)-stellarChecksumLen:]

	crcNum := crc.CalculateCRC(crc.XMODEM, payload)
	expectedChecksum := []byte{0, 0}
	binary.LittleEndian.PutUint16(expectedChecksum, uint16(crcNum))

	if bytes.Compare(expectedChecksum, checksum) == 0 {
		return true
	}

	return false
}
