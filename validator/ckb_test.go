package validator

import (
	"reflect"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCKBValidateAddress(t *testing.T) {
	validator := &CKB{}

	var mainnetCases = map[string]*Result{
		"ckb1qyq0fwdxwjehj30vuymzcypx55zpqm2uw79ql0hyu8": {true, CKBShortPayload, ""},
		"ckb1qyqpk0apd93llee89ldxn7fv4xzyrzqwskws98xs26": {true, CKBShortPayload, ""},
		"ckb1qyqx8emxzme2rcdxluuxmv44pkvhx5ev27jqk0sqpj": {true, CKBShortPayload, ""},
		"ckb1qyqt8xaupvm8837nv3gtc9x0ekkj64vud3jqfwyw5v": {true, CKBShortPayload, ""},
		"ckb1qyq5lv479ewscx3ms620sv34pgeuz6zagaaqklhtgg": {true, CKBShortPayload, ""},
	}

	for addr, result := range mainnetCases {
		assert.True(t, reflect.DeepEqual(validator.ValidateAddress(addr, Mainnet), result), addr)
		assert.False(t, reflect.DeepEqual(validator.ValidateAddress(addr, Testnet), result), addr)
	}

	var testnetCases = map[string]*Result{
		"ckt1qyqrdsefa43s6m882pcj53m4gdnj4k440axqswmu83": {true, CKBShortPayload, ""},
		"ckt1qyqpgc5cjcw4e45ahpn55s0yqhappes0cxvsgxtluv": {true, CKBShortPayload, ""},
		"ckt1qyq2tzypn6hfnkn2xtm5fa3s8ej83q357evqe67uze": {true, CKBShortPayload, ""},
		"ckt1qyqlqn8vsj7r0a5rvya76tey9jd2rdnca8lqh4kcuq": {true, CKBShortPayload, ""},
	}

	for addr, result := range testnetCases {
		assert.False(t, reflect.DeepEqual(validator.ValidateAddress(addr, Mainnet), result), addr)
		assert.True(t, reflect.DeepEqual(validator.ValidateAddress(addr, Testnet), result), addr)
	}

	var invalidCases = map[string]*Result{
		"cbt1qyqrdsefa43s6m882pcj53m4gdnj4k440axqswmu83": {false, Unknown, ""},
		"ckb1qyqt8xaupvm8837nv3gtc9x0ekkj64vud3jqfwyw4v": {false, Unknown, ""},
		"ckb1qyqpk0apd93llee89ldxn7fv4xzyrzqwskws98s26":  {false, Unknown, ""},
		"NBVH-33D3BQ2TWUR6TJNZTZOVR2M3A-ZDUZ-RDRK-R4EA":  {false, Unknown, ""},
		// # valid full payload data address, but we don't support yet
		"ckb1qsvf96jqmq4483ncl7yrzfzshwchu9jd0glq4yy5r2jcsw04d7xlydkr98kkxrtvuag8z2j8w4pkw2k6k4l5czfy37k": {false, Unknown, ""},
		"ckt1q2n9dutjk669cfznq7httfar0gtk7qp0du3wjfvzck9l0w3k9eqhvdkr98kkxrtvuag8z2j8w4pkw2k6k4l5czshhac": {false, Unknown, ""},
	}

	for addr, result := range invalidCases {
		assert.True(t, reflect.DeepEqual(validator.ValidateAddress(addr, Mainnet), result), addr)
		assert.True(t, reflect.DeepEqual(validator.ValidateAddress(addr, Testnet), result), addr)
	}
}
