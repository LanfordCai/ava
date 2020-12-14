package crypto

import (
	"crypto/sha256"

	"github.com/ethereum/go-ethereum/crypto/sha3"
	"github.com/minio/blake2b-simd"
	"golang.org/x/crypto/ripemd160"
)

// Sha256 ...
func Sha256(data []byte) []byte {
	h := sha256.New()
	h.Write(data)
	return h.Sum(nil)
}

// DoubleSha256 ...
func DoubleSha256(data []byte) []byte {
	return Sha256(Sha256(data))
}

// Ripemd160 ...
func Ripemd160(data []byte) []byte {
	h := ripemd160.New()
	h.Write(data)
	return h.Sum(nil)
}

// Blake2b256 ...
func Blake2b256(data []byte) []byte {
	h := blake2b.New256()
	h.Write(data)
	return h.Sum(nil)
}

// Blake2b512 ...
func Blake2b512(data []byte) []byte {
	h := blake2b.New512()
	h.Write(data)
	return h.Sum(nil)
}

// Keccak256 ...
func Keccak256(data []byte) []byte {
	h := sha3.NewKeccak256()
	h.Write(data)
	return h.Sum(nil)
}

// Blake2bKeccak256 ...
func Blake2bKeccak256(data []byte) []byte {
	return Keccak256(Blake2b256(data))
}
