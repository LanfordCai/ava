package validator

import "github.com/btcsuite/btcutil/bech32"

// CKB ...
type CKB struct{}

var _ Validator = (*CKB)(nil)

const ckbShortPayloadFormatType = 1

const ckbBlake2bCodeHashType = 0
const ckbMultisigHashType = 1

// ValidateAddress returns validate result of ckb address
// only support short payload format, see:
// https://github.com/nervosnetwork/rfcs/blob/master/rfcs/0021-ckb-address-format/0021-ckb-address-format.md
func (v *CKB) ValidateAddress(addr string, network NetworkType) *Result {
	hrp, data, err := bech32.Decode(addr)
	if hrp != v.AddressHrp(network) || err != nil {
		return &Result{Success, false, Unknown, ""}
	}

	program, err := bech32.ConvertBits(data, 5, 8, true)
	if err != nil {
		return &Result{Success, false, Unknown, ""}
	}

	addrType := program[0]
	codeHashIdx := program[1]
	if addrType == ckbShortPayloadFormatType &&
		(codeHashIdx == ckbBlake2bCodeHashType || codeHashIdx == ckbMultisigHashType) {
		return &Result{Success, true, CKBShortPayload, ""}
	}

	return &Result{Success, false, Unknown, ""}
}

// AddressHrp returns hrps of ckb according to the network
func (v *CKB) AddressHrp(network NetworkType) string {
	if network == Mainnet {
		return "ckb"
	}
	return "ckt"
}
