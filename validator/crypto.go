package validator

import (
	"github.com/btcsuite/btcutil/bech32"
	"github.com/ethereum/go-ethereum/crypto/sha3"
	"github.com/minio/blake2b-simd"
)

func keccak256Sum(data []byte) []byte {
	hash := sha3.NewKeccak256()
	var buf []byte
	hash.Write(data)
	buf = hash.Sum(buf)

	return buf
}

func blake2b256Sum(data []byte) []byte {
	hash := blake2b.Sum256(data)
	return hash[:]
}

func bech32AddrDecode(address string) (string, []byte, error) {
	hrp, data, err := bech32.Decode(address)
	if err != nil || len(data) < 1 {
		return "", nil, err
	}
	program, err := bech32.ConvertBits(data, 5, 8, false)
	if err != nil {
		return "", nil, err
	}
	return hrp, program, nil
}

func segwitAddrDecode(address string) (string, byte, []byte, error) {
	hrp, data, err := bech32.Decode(address)
	if err != nil || len(data) < 1 {
		return "", 255, nil, err
	}
	version := data[0]
	program, err := bech32.ConvertBits(data[1:], 5, 8, false)
	if err != nil {
		return "", 255, nil, err
	}
	return hrp, version, program, nil
}
