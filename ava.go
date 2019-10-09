package ava

import (
	"fmt"
	"os"
	"strings"

	"github.com/LanfordCai/ava/internal/common"
	"github.com/LanfordCai/ava/internal/utils"
	"github.com/LanfordCai/ava/pkg/aeternity"
	"github.com/LanfordCai/ava/pkg/bitcoin"
	"github.com/LanfordCai/ava/pkg/bitcoingold"
	"github.com/LanfordCai/ava/pkg/bytom"
	"github.com/LanfordCai/ava/pkg/dash"
	"github.com/LanfordCai/ava/pkg/dogecoin"
	"github.com/LanfordCai/ava/pkg/eos"
	"github.com/LanfordCai/ava/pkg/ethereum"
	"github.com/LanfordCai/ava/pkg/litecoin"
	"github.com/LanfordCai/ava/pkg/neo"
	"github.com/LanfordCai/ava/pkg/qtum"
	"github.com/LanfordCai/ava/pkg/ripple"
	"github.com/LanfordCai/ava/pkg/sia"
	"github.com/LanfordCai/ava/pkg/zcash"
)

// Validator - Address validator
type Validator interface {
	ValidateAddress(addr string, isTestnet bool) (isValid bool, msg string)
}

// NewValidator - Create address validator according to given chain name
func NewValidator(chainName string) (Validator, error) {
	chainName = strings.ToLower(chainName)
	switch chainName {
	case "bitcoin", "bitcoincash", "bitcoindiamond", "superbitcoin":
		types := getEnalbedTypes(chainName)
		return bitcoin.New(types)
	case "litecoin":
		types := getEnalbedTypes(chainName)
		return litecoin.New(types)
	case "bytom":
		types := getEnalbedTypes(chainName)
		return bytom.New(types)
	case "dogecoin":
		types := getEnalbedTypes(chainName)
		return dogecoin.New(types)
	case "zcash":
		types := getEnalbedTypes(chainName)
		return zcash.New(types)
	case "sia", "siacore", "siaclassic":
		return sia.New(), nil
	case "qtum":
		types := getEnalbedTypes(chainName)
		return qtum.New(types)
	case "neo":
		return neo.New(), nil
	case "ripple":
		return ripple.New(), nil
	case "ethereum", "ethereumclassic", "gojoychain", "wanchain", "vechain":
		return ethereum.New(), nil
	case "dash":
		types := getEnalbedTypes(chainName)
		return dash.New(types)
	case "eos":
		baseURL := os.Getenv("AVA_EOS_BASE_URL")
		whitelist := getContractWhiteList(chainName)
		return eos.New(baseURL, whitelist), nil
	case "bitcoingold":
		types := getEnalbedTypes(chainName)
		return bitcoingold.New(types)
	case "aeternity":
		return aeternity.New(), nil
	default:
		return nil, common.ErrUnsupportedChain
	}
}

func getEnalbedTypes(chainName string) []string {
	chainName = strings.ToUpper(chainName)
	typesStr := os.Getenv(fmt.Sprintf("AVA_%s_ENABLED_ADDR_TYPES", chainName))
	return utils.SplitWithComma(typesStr)
}

func getContractWhiteList(chainName string) []string {
	chainName = strings.ToUpper(chainName)
	whitelistStr := os.Getenv(fmt.Sprintf("AVA_%s_CONTRACT_WHITELIST", chainName))
	return utils.SplitWithComma(whitelistStr)
}
