package validator

import (
	"bytes"

	"github.com/LanfordCai/ava/base58check"
)

// Tezos ...
type Tezos struct{}

var _ Validator = (*Tezos)(nil)

var tezosPrefixes = [][]byte{
	{6, 161, 159},
	{6, 161, 161},
	{6, 161, 164},
}

// ValidateAddress returns validate result of tezos address
// NOTE: don't support delegated address(KT1 address)
func (v *Tezos) ValidateAddress(addr string, network NetworkType) *Result {
	data, err := base58check.BitcoinEncoder.CheckDecode(addr)
	if err != nil {
		return &Result{Success, false, Unknown, ""}
	}

	prefix := []byte{data[0], data[1], data[2]}
	if contains(tezosPrefixes, prefix) {
		return &Result{Success, true, Normal, ""}
	}

	return &Result{Success, false, Unknown, ""}
}

func contains(s [][]byte, b []byte) bool {
	for _, r := range s {
		if bytes.Compare(r, b) == 0 {
			return true
		}
	}
	return false
}
