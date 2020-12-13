package cashaddr

import (
	"bytes"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestExpandPrefix(t *testing.T) {
	expanded := expandPrefix("bitcoincash")
	expected := []byte{2, 9, 20, 3, 15, 9, 14, 3, 1, 19, 8, 0}
	assert.True(t, bytes.Compare(expanded, expected) == 0)
}

func TestPolymod(t *testing.T) {
	data := []byte{
		0, 0, 0, 0, 0, 0, 0, 0, 0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18,
		19, 20, 21, 22, 23, 24, 25, 26, 27, 28, 29, 30, 31, 0, 0, 0, 0, 0, 0, 0, 0,
	}

	expected := 0x724AFE7AF2

	assert.True(t, polymod(data) == expected)
}

func TestToLegacyAddr(t *testing.T) {
	cases := map[string]string{
		"bitcoincash:qzy9dedpwm53cqgcr2e3z9qy788cca6y0ce8mfkezh": "1DRtzoHTzEkj2ayMU1qyVjFRRUsyx7CacZ",
		"bitcoincash:pp297ralj957hvzyyvgmfe5vtz5c2szxqcehz8lc7f": "39P8ZeWRGVLsH4smmsjLM1rybDmMLrUKRq",
		"bitcoincash:ppm2qsznhks23z7629mms6s4cwef74vcwvn0h829pq": "3CWFddi6m4ndiGyKqzYvsFYagqDLPVMTzC",
		"bchtest:qzjncxjjzga5mr6qyy7al7g46zavnar7ggmc50uj8l":     "mvadnisqDQP3pzxWxtPzkdhfAgBGLThNEZ",
	}

	for k, v := range cases {
		got, err := ToLegacyAddr(k)
		assert.NoError(t, err)
		assert.Equal(t, got, v)
	}
}

func TestFromLegacyAddr(t *testing.T) {
	cases := map[string]string{
		"1BpEi6DfDAUFd7GtittLSdBeYJvcoaVggu": "bitcoincash:qpm2qsznhks23z7629mms6s4cwef74vcwvy22gdx6a",
		"mkAb1V4jYomRmhqpek5ZpyrcZVscZmEGEf": "bchtest:qqe0az5j8gwk2uts9v2ywdv6sd32uut9hq9hc92pee",
		"3CWFddi6m4ndiGyKqzYvsFYagqDLPVMTzC": "bitcoincash:ppm2qsznhks23z7629mms6s4cwef74vcwvn0h829pq",
		"39P8ZeWRGVLsH4smmsjLM1rybDmMLrUKRq": "bitcoincash:pp297ralj957hvzyyvgmfe5vtz5c2szxqcehz8lc7f",
	}

	for k, v := range cases {
		got, err := FromLegacyAddr(k)
		assert.NoError(t, err)
		assert.Equal(t, got, v)
	}
}

// assert Converter.from_legacy("1BpEi6DfDAUFd7GtittLSdBeYJvcoaVggu") ==
// "bitcoincash:qpm2qsznhks23z7629mms6s4cwef74vcwvy22gdx6a"

// assert Converter.from_legacy("mkAb1V4jYomRmhqpek5ZpyrcZVscZmEGEf") ==
// "bchtest:qqe0az5j8gwk2uts9v2ywdv6sd32uut9hq9hc92pee"

// assert Converter.from_legacy("3CWFddi6m4ndiGyKqzYvsFYagqDLPVMTzC") ==
// "bitcoincash:ppm2qsznhks23z7629mms6s4cwef74vcwvn0h829pq"

// assert Converter.from_legacy("39P8ZeWRGVLsH4smmsjLM1rybDmMLrUKRq") ==
// "bitcoincash:pp297ralj957hvzyyvgmfe5vtz5c2szxqcehz8lc7f"
