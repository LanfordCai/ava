package validator

// NetworkType ...
type NetworkType string

// List of NetworkType
const (
	Mainnet NetworkType = "Mainnet"
	Testnet NetworkType = "Testnet"
)

// AddressType ...
type AddressType string

// List AddressType
const (
	Unknown            AddressType = "Unknown"
	Normal             AddressType = "Normal"
	P2PKH              AddressType = "P2PKH"
	P2SH               AddressType = "P2SH"
	P2WPKH             AddressType = "P2WPKH"
	P2WSH              AddressType = "P2WSH"
	CashAddrP2PKH      AddressType = "CashAddrP2PKH"
	CashAddrP2SH       AddressType = "CashAddrP2SH"
	HexWithChecksum    AddressType = "HexWithChecksum"
	HexWithoutChecksum AddressType = "HexWithoutChecksum"
)

// Result is the validate result of address
type Result struct {
	IsValid bool
	Type    AddressType
	Msg     string
}
