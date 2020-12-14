package base58check

// ChecksumType ...
type ChecksumType int

// List of ChecksumType
const (
	ChecksumDoubleSha256 ChecksumType = iota
	ChecksumSha256
	ChecksumRipemd160
	ChecksumBlake2bKeccak256
)
