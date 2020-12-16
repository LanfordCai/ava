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

// List of AddressType
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
	CKBShortPayload    AddressType = "CKBShortPayload"
	FilID              AddressType = "FilID"
	FilActor           AddressType = "FilActor"
	FilSecp256k1       AddressType = "FileSecp256k1"
	FilBLS             AddressType = "FileBLS"
)

// ValidateStatus ...
type ValidateStatus string

// List of ValidateStatus
const (
	Success ValidateStatus = "Success"
	Failure ValidateStatus = "Failure"
)

// Result is the validate result of address
type Result struct {
	Status  ValidateStatus
	IsValid bool
	Type    AddressType
	Msg     string
}
