package ava

import (
	"fmt"
	"os"
	"strings"

	"github.com/LanfordCai/ava/internal/common"
	"github.com/LanfordCai/ava/pkg/bitcoin"
	"github.com/LanfordCai/ava/pkg/bytom"
	"github.com/LanfordCai/ava/pkg/dash"
	"github.com/LanfordCai/ava/pkg/dogecoin"
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
	ValidateAddress(addr string, isTestnet bool) common.ValidationResult
}

// NewValidator - Create address validator according to given chain name
func NewValidator(chainName string) (Validator, error) {
	chainName = strings.ToLower(chainName)
	switch chainName {
	case "bitcoin", "bitcoincash", "bitcoindiamond":
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
	default:
		return nil, common.ErrUnsupportedChain
	}
}

func getEnalbedTypes(chainName string) []string {
	chainName = strings.ToUpper(chainName)
	key := fmt.Sprintf("AVA_%s_ENABLED_ADDR_TYPES", chainName)
	typesStr := os.Getenv(key)
	if len(typesStr) == 0 {
		return []string{}
	}
	types := strings.Split(typesStr, ",")
	for i := 0; i < len(types); i++ {
		types[i] = strings.Trim(types[i], " ")
	}

	return types
}
