package base58check

import (
	"ava/crypto"
	"bytes"
	"math/big"
	"strings"
)

// BitcoinEncoder ...
var BitcoinEncoder = &Encoder{
	enc:          Encoding{alphabet: "123456789ABCDEFGHJKLMNPQRSTUVWXYZabcdefghijkmnopqrstuvwxyz"},
	ChecksumLen:  4,
	ChecksumType: ChecksumDoubleSha256,
}

// RippleEncoder ...
var RippleEncoder = &Encoder{
	enc:          Encoding{alphabet: "rpshnaf39wBUDNEGHJKLM4PQRST7VWXYZ2bcdeCg65jkm8oFqi1tuvAxyz"},
	ChecksumLen:  4,
	ChecksumType: ChecksumDoubleSha256,
}

var bigRadix = big.NewInt(58)
var bigZero = big.NewInt(0)

// Encoder ...
type Encoder struct {
	enc          Encoding
	ChecksumLen  int
	ChecksumType ChecksumType
}

// CheckEncode ...
func (e *Encoder) CheckEncode(data []byte) string {
	checksum := e.calcChecksum(data)
	payload := make([]byte, len(data))
	copy(payload[:], data[:])
	payload = append(payload, checksum...)
	return e.Encode(payload)
}

// CheckDecode ...
func (e *Encoder) CheckDecode(str string) ([]byte, error) {
	decoded := e.Decode(str)
	if len(decoded) < e.ChecksumLen {
		return nil, ErrInvalidInput
	}
	payload, err := e.validateChecksum(decoded)
	if err != nil {
		return nil, err
	}
	return payload, nil
}

// Encode ...
// Copyright (c) 2013-2015 The btcsuite developers
// Use of this source code is governed by an ISC
// license that can be found in the LICENSE file.
// Modified by LanfordCai
func (e *Encoder) Encode(b []byte) string {
	x := new(big.Int)
	x.SetBytes(b)
	alphabetIdx0 := e.enc.alphabet[0]

	answer := make([]byte, 0, len(b)*136/100)
	for x.Cmp(bigZero) > 0 {
		mod := new(big.Int)
		x.DivMod(x, bigRadix, mod)
		answer = append(answer, e.enc.alphabet[mod.Int64()])
	}

	// leading zero bytes
	for _, i := range b {
		if i != 0 {
			break
		}
		answer = append(answer, alphabetIdx0)
	}

	// reverse
	alen := len(answer)
	for i := 0; i < alen/2; i++ {
		answer[i], answer[alen-1-i] = answer[alen-1-i], answer[i]
	}

	return string(answer)
}

// Decode ...
// Copyright (c) 2013-2015 The btcsuite developers
// Use of this source code is governed by an ISC
// license that can be found in the LICENSE file.
// Modified by LanfordCai
func (e *Encoder) Decode(b string) []byte {
	answer := big.NewInt(0)
	j := big.NewInt(1)
	alphabetIdx0 := e.enc.alphabet[0]

	set := e.enc.codeToNum()
	scratch := new(big.Int)
	for i := len(b) - 1; i >= 0; i-- {
		tmp := set[rune(b[i])]
		if tmp == 255 {
			return []byte("")
		}
		scratch.SetInt64(int64(tmp))
		scratch.Mul(j, scratch)
		answer.Add(answer, scratch)
		j.Mul(j, bigRadix)
	}

	tmpval := answer.Bytes()

	var numZeros int
	for numZeros = 0; numZeros < len(b); numZeros++ {
		if b[numZeros] != alphabetIdx0 {
			break
		}
	}
	flen := numZeros + len(tmpval)
	val := make([]byte, flen)
	copy(val[numZeros:], tmpval)

	return val
}

func (e *Encoder) verifyBase(str string) error {
	for _, r := range str {
		if !strings.Contains(e.enc.alphabet, string(r)) {
			return ErrInvalidInput
		}
	}
	return nil
}

func (e *Encoder) validateChecksum(data []byte) ([]byte, error) {
	payload := data[:len(data)-e.ChecksumLen]
	checksum := data[len(data)-e.ChecksumLen:]

	expectedChecksum := e.calcChecksum(payload)
	if bytes.Compare(checksum, expectedChecksum) == 0 {
		return payload, nil
	}
	return nil, ErrInvalidChecksum
}

func (e *Encoder) calcChecksum(payload []byte) []byte {
	switch e.ChecksumType {
	case ChecksumDoubleSha256:
		return crypto.DoubleSha256(payload)[:e.ChecksumLen]
	case ChecksumSha256:
		return crypto.Sha256(payload)[:e.ChecksumLen]
	case ChecksumRipemd160:
		return crypto.Ripemd160(payload)[:e.ChecksumLen]
	case ChecksumBlake2bKeccak256:
		return crypto.Blake2bKeccak256(payload)[:e.ChecksumLen]
	default:
		panic(ErrUnsupportedChecksumType)
	}
}

// Encoding ...
type Encoding struct {
	alphabet string
}

func (e *Encoding) codeToNum() map[rune]byte {
	result := make(map[rune]byte)
	for i, r := range e.alphabet {
		result[r] = byte(i)
	}
	return result
}
