package validator

import (
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
