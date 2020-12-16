package utils

import (
	"fmt"
	"os"
	"strings"

	"github.com/LanfordCai/ava/validator"
)

// IsUnsupportedAddressType ...
func IsUnsupportedAddressType(chain string, addrType validator.AddressType) bool {
	chainName := strings.ToUpper(chain)
	key := fmt.Sprintf("AVA_%s_UNSUPPORTED_ADDR_TYPES", chainName)
	value := os.Getenv(key)
	if value == "" {
		return false
	}

	types := normalizeListString(value)
	return Contains(types, strings.ToLower(string(addrType)))
}

// AddressInContractWhiteList ...
func AddressInContractWhiteList(chain, addr string) bool {
	chainName := strings.ToUpper(chain)
	key := fmt.Sprintf("AVA_%s_CONTRACT_WHITELIST", chainName)
	value := os.Getenv(key)
	if key == "" {
		return false
	}

	addrs := normalizeListString(value)
	return Contains(addrs, strings.ToLower(addr))
}

// EndpointForChain ...
func EndpointForChain(chain string) string {
	chainName := strings.ToUpper(chain)
	key := fmt.Sprintf("AVA_%s_ENDPOINT", chainName)
	return strings.TrimSpace(os.Getenv(key))
}

// NeedCheckContractAddress ...
func NeedCheckContractAddress(chain string) bool {
	value := os.Getenv("AVA_CHAINS_NEED_CHECK_CONTRACT_ADDR")
	if value == "" {
		return false
	}

	chains := normalizeListString(value)
	return Contains(chains, strings.ToLower(chain))
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
	case "eos", "yas", "wax", "yottachain", "abbc":
		client := validator.EOSClient{Endpoint: EndpointForChain(chain)}
		return &validator.EOS{Client: &client}
	case "bitshares", "gxshares":
		client := validator.BitsharesClient{Endpoint: EndpointForChain(chain)}
		return &validator.Bitshares{Client: &client}
	case "iost":
		client := validator.IOSTClient{Endpoint: EndpointForChain(chain)}
		return &validator.IOST{Client: &client}
	case "stellar":
		client := validator.StellarClient{Endpoint: EndpointForChain(chain)}
		return &validator.Stellar{Client: &client}
	case "tera":
		client := validator.TeraClient{Endpoint: EndpointForChain(chain)}
		return &validator.Tera{Client: &client}
	case "oasis":
		client := validator.RosettaClient{Endpoint: EndpointForChain(chain)}
		return &validator.Oasis{Client: &client}
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
	default:
		return nil
	}
}

func normalizeListString(listString string) []string {
	items := strings.Split(strings.ToLower(listString), ",")
	normalizedItems := []string{}
	for _, i := range items {
		normalizedItems = append(normalizedItems, strings.TrimSpace(i))
	}
	return normalizedItems
}
