package utils

import (
	"fmt"
	"os"
	"strings"

	"github.com/LanfordCai/ava/validator"
)

// GetUnsupportedAddressTypes ...
func GetUnsupportedAddressTypes(chain string) []string {
	chainName := strings.ToUpper(chain)
	key := fmt.Sprintf("AVA_%s_UNSUPPORTED_ADDR_TYPES", chainName)
	value := os.Getenv(key)
	if value == "" {
		return []string{}
	}

	types := strings.Split(value, ",")
	normalizedTypes := []string{}
	for _, t := range types {
		normalizedTypes = append(normalizedTypes, strings.TrimSpace(t))
	}
	return normalizedTypes
}

// Contains ...
func Contains(strs []string, str string) bool {
	for _, s := range strs {
		if s == str {
			return true
		}
	}
	return false
}

// GetValidatorByChain ...
func GetValidatorByChain(chain string) validator.Validator {
	switch chainName := strings.ToLower(chain); chainName {
	case "bitcoin", "bitcoinsegwit", "bitcoindiamond", "superbitcoin":
		return &validator.Bitcoin{}
	case "ethereum", "ethereumclassic", "wanchain", "fusion", "vechain", "bittrustsystem", "klaytn":
		return &validator.Ethereum{}
	case "neo", "ontology":
		return &validator.NEO{}
	case "bitcoincash", "bitcoinabc":
		return &validator.BitcoinCash{}
	case "litecoin":
		return &validator.Litecoin{}
	case "zcash":
		return &validator.Zcash{}
	case "qtum":
		return &validator.Qtum{}
	case "sia", "siaclassic", "siacore":
		return &validator.Sia{}
	case "ripple":
		return &validator.Ripple{}
	case "bytom":
		return &validator.Bytom{}
	case "bitcoinsv":
		return &validator.BitcoinSV{}
	case "dogecoin":
		return &validator.Dogecoin{}
	case "bitcoingold":
		return &validator.BitcoinGold{}
	case "tron":
		return &validator.Tron{}
	case "cosmos":
		return &validator.Cosmos{}
	case "dash":
		return &validator.Dash{}
	case "aeternity":
		return &validator.Aeternity{}
	case "tezos":
		return &validator.Tezos{}
	case "ycash":
		return &validator.Ycash{}
	case "chainx":
		return &validator.ChainX{}
	case "zilliqa":
		return &validator.Zilliqa{}
	case "iris":
		return &validator.Iris{}
	case "vsystems":
		return &validator.Vsystems{}
	case "zeepin":
		return &validator.Zeepin{}
	case "ckb":
		return &validator.CKB{}
	case "classzz":
		return &validator.Classzz{}
	case "lbry":
		return &validator.Lbry{}
	case "handshake":
		return &validator.Handshake{}
	case "ucacoin":
		return &validator.Ucacoin{}
	case "polkadot":
		return &validator.Polkadot{}
	case "wayfcoin":
		return &validator.Wayfcoin{}
	case "arweave":
		return &validator.Arweave{}
	case "oasis":
		return &validator.Oasis{}
	default:
		return nil
	}
}
