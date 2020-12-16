package cashaddr

import (
	"bytes"
	"fmt"
	"strings"

	"github.com/btcsuite/btcutil/base58"
	"github.com/btcsuite/btcutil/bech32"
)

var enc = newBase32Encoder("qpzry9x8gf2tvdw0s3jn54khce6mua7l")

// FromLegacyAddr ...
func FromLegacyAddr(addr string) (string, error) {
	_, version, err := base58.CheckDecode(addr)
	if err != nil {
		return "", ErrInvalidCashAddr
	}
	prefix := calcPrefix(version)
	packedAddr := packAddress(addr)

	payload := make([]byte, len(packedAddr))
	copy(payload[:], packedAddr[:])
	payload = append(payload, []byte{0, 0, 0, 0, 0, 0, 0, 0}...)
	checksum := calcChecksum(prefix, payload)

	body := enc.encode(append(packedAddr, checksum...))
	return fmt.Sprintf("%s:%s", prefix, string(body)), nil
}

// ToLegacyAddr ...
func ToLegacyAddr(addr string) (string, error) {
	components := strings.Split(addr, ":")
	if len(components) != 2 {
		return "", ErrInvalidCashAddr
	}

	prefix := components[0]
	data := components[1]

	if len(data) < 8 {
		return "", ErrInvalidCashAddr
	}

	decodedData := enc.decode(data)
	packedAddrBytes := decodedData[:len(decodedData)-8]
	checksumBytes := decodedData[len(decodedData)-8:]

	payload := make([]byte, len(packedAddrBytes))
	copy(payload[:], packedAddrBytes[:])
	payload = append(payload, []byte{0, 0, 0, 0, 0, 0, 0, 0}...)
	expectedChecksum := calcChecksum(prefix, payload)

	if bytes.Compare(expectedChecksum, checksumBytes) != 0 {
		return "", ErrInvalidCashAddr
	}

	packedAddr := data[:len(data)-8]
	return unpackAddress(packedAddr, prefix)
}

func unpackAddress(packedAddr string, prefix string) (string, error) {
	packedAddrBytes := enc.decode(packedAddr)

	data, err := bech32.ConvertBits(packedAddrBytes, 5, 8, false)
	if err != nil {
		return "b", ErrInvalidCashAddr
	}

	prefixByte := byte(0)
	if prefix == "bitcoincash" && packedAddr[0] == 'q' {
		prefixByte = 0
	} else if prefix == "bitcoincash" && packedAddr[0] == 'p' {
		prefixByte = 5
	} else if prefix == "bchtest" && packedAddr[0] == 'q' {
		prefixByte = 111
	} else if prefix == "bchtest" && packedAddr[0] == 'p' {
		prefixByte = 196
	} else if prefix == "classzz" && packedAddr[0] == 'c' {
		prefixByte = 192
	}

	return base58.CheckEncode(data[1:], prefixByte), nil
}

func packAddress(addr string) []byte {
	decoded, version, _ := base58.CheckDecode(addr)
	encodedSize := 0
	switch l := len(decoded) * 8; l {
	case 160:
		encodedSize = 0
	case 192:
		encodedSize = 1
	case 224:
		encodedSize = 2
	case 256:
		encodedSize = 3
	case 320:
		encodedSize = 4
	case 448:
		encodedSize = 6
	case 512:
		encodedSize = 7
	default:
		panic(ErrInvalidLegacyAddr)
	}

	addrType := calcAddrType(version)
	versionByte := addrType<<3 | byte(encodedSize)
	result, err := bech32.ConvertBits(append([]byte{versionByte}, decoded...), 8, 5, true)
	if err != nil {
		panic(ErrInvalidLegacyAddr)
	}
	return result
}

func calcChecksum(prefix string, packedAddr []byte) []byte {
	expandedPrefix := expandPrefix(prefix)
	poly := polymod(append(expandedPrefix, packedAddr...))
	result := []byte{}
	for i := 0; i < 8; i++ {
		result = append(result, byte(poly>>(5*(7-i))&0x1F))
	}
	return result
}

func calcPrefix(version byte) string {
	switch {
	case version == 0 || version == 5:
		return "bitcoincash"
	case version == 111 || version == 196:
		return "bchtest"
	case version == 192:
		return "classzz"
	default:
		panic(ErrUnsupported.Error())
	}
}

func calcAddrType(version byte) byte {
	switch {
	case version == 0 || version == 111:
		return 0
	case version == 5 || version == 196:
		return 1
	case version == 192:
		return 24
	default:
		panic(ErrUnsupported.Error())
	}
}

func expandPrefix(prefix string) []byte {
	result := []byte{}
	prefixData := []byte(prefix)

	for _, b := range prefixData {
		result = append(result, b&0x1F)
	}

	return append(result, 0x00)
}

func polymod(data []byte) int {
	acc := 1
	for _, b := range data {
		c0 := acc >> 35
		acc = ((acc & 0x07FFFFFFFF) << 5) ^ int(b)

		if (c0 & 0x01) != 0 {
			acc ^= 0x98F2BC8E61
		}

		if (c0 & 0x02) != 0 {
			acc ^= 0x79B76D99E2
		}

		if (c0 & 0x04) != 0 {
			acc ^= 0xF33E5FB3C4
		}

		if (c0 & 0x08) != 0 {
			acc ^= 0xAE2EABE2A8
		}

		if (c0 & 0x10) != 0 {
			acc ^= 0x1E4F43E470
		}
	}

	return acc ^ 1
}
