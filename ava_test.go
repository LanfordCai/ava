package ava

import (
	"testing"
)

var validAddresses = []Address{
	Address{"Bitcoin", "1HX2swQNH9ezE8RagPPAaSBTcK3in9xWYF", false},
	Address{"Litecoin", "LUd28DgwFmae8thFhaPYHGVmps4XockqfQ", false},
	Address{"Ethereum", "0x001d3f1ef827552ae1114027bd3ecf1f086ba0f9", false},
	Address{"EthereumClassic", "0x001d3f1ef827552ae1114027bd3ecf1f086ba0f9", false},
	Address{"EOS", "pxneosincome", false},
	Address{"Neo", "AVKeLwu1uRtM7Lf7ckkqrbtkvss7jkN38N", false},
	Address{"Qtum", "QhzgLgC9DYxn9sTXWfcsFdKbFbbx2DKDL4", false},
	Address{"Ripple", "rPFLkxQk6xUGdGYEykqe7PR25Gr7mLHDc8", false},
	Address{"SiaCore", "00322de65dca5d1683989d5f1647261393d6faf98c0e61dbb098bc9a21473995fba12fcabd9f", false},
	Address{"Zcash", "t1J1YojHoLb5L9pvRi1sCSNgyReHrK4EbDw", false},
	Address{"Bytom", "bm1qw508d6qejxtdg4y5r3zarvary0c5xw7k23gyyf", false},
}

var invalidAddresses = []Address{
	Address{"Bitcoin", "0HX2swQNH9ezE8RagPPAaSBTcK3in9xWYF", false},
	Address{"Litecoin", "LU328DgwFmae8thFhaPYHGVmps4XockqfQ", false},
	Address{"Ethereum", "0x!01d3f1ef827552ae1114027bd3ecf1f086ba0f8", false},
	Address{"EOS", "pxneosincome123", false},
	Address{"Neo", "AVKeLwu1uRtM7L07ckkqrbtkvss7jkN38N", false},
	Address{"Qtum", "QhzgLgC9DYxn9ETXWfcsFdKbFbbx2DKDL4", false},
	Address{"Ripple", "1PFLkxQk6xUGdGYEykqe7PR25Gr7mLHDc8", false},
	Address{"SiaCore", "005d1683989d5f1647261393d6faf98c0e61dbb098bc9a21473995fba12fcabd9f", false},
	Address{"Zcash", "t1J13ojHoLb5L9pvRi1sCSNgyReHrK4EbDw", false},
	Address{"Bytom", "bm2qw508d6qejxtdg4y5r3zarvary0c5xw7k23gyyf", false},
}

func TestIsValid(t *testing.T) {
	for _, a := range validAddresses {
		if !a.IsValid() {
			t.Errorf("%s check result is wrong!", a.Address)
		}
	}

	for _, a := range invalidAddresses {
		if a.IsValid() {
			t.Errorf("%s check result is wrong!", a.Address)
		}
	}
}
