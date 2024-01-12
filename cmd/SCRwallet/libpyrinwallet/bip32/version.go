package bip32

import "github.com/pkg/errors"

// BitcoinMainnetPrivate is the version that is used for
// bitcoin mainnet bip32 private extended keys.
// Ecnodes to xprv in base58.
var BitcoinMainnetPrivate = [4]byte{
	0x04,
	0x88,
	0xad,
	0xe4,
}

// BitcoinMainnetPublic is the version that is used for
// bitcoin mainnet bip32 public extended keys.
// Ecnodes to xpub in base58.
var BitcoinMainnetPublic = [4]byte{
	0x04,
	0x88,
	0xb2,
	0x1e,
}

// SCRMainnetPrivate is the version that is used for
// SCR mainnet bip32 private extended keys.
// Ecnodes to xprv in base58.
var SCRMainnetPrivate = [4]byte{
	0x03,
	0x8f,
	0x2e,
	0xf4,
}

// SCRMainnetPublic is the version that is used for
// SCR mainnet bip32 public extended keys.
// Ecnodes to kpub in base58.
var SCRMainnetPublic = [4]byte{
	0x03,
	0x8f,
	0x33,
	0x2e,
}

// SCRTestnetPrivate is the version that is used for
// SCR testnet bip32 public extended keys.
// Ecnodes to ktrv in base58.
var SCRTestnetPrivate = [4]byte{
	0x03,
	0x90,
	0x9e,
	0x07,
}

// SCRTestnetPublic is the version that is used for
// SCR testnet bip32 public extended keys.
// Ecnodes to ktub in base58.
var SCRTestnetPublic = [4]byte{
	0x03,
	0x90,
	0xa2,
	0x41,
}

// SCRpadevnetPrivate is the version that is used for
// SCR devnet bip32 public extended keys.
// Ecnodes to kdrv in base58.
var SCRpadevnetPrivate = [4]byte{
	0x03,
	0x8b,
	0x3d,
	0x80,
}

// SCRpadevnetPublic is the version that is used for
// SCR devnet bip32 public extended keys.
// Ecnodes to xdub in base58.
var SCRpadevnetPublic = [4]byte{
	0x03,
	0x8b,
	0x41,
	0xba,
}

// SCRSimnetPrivate is the version that is used for
// SCR simnet bip32 public extended keys.
// Ecnodes to ksrv in base58.
var SCRSimnetPrivate = [4]byte{
	0x03,
	0x90,
	0x42,
	0x42,
}

// SCRSimnetPublic is the version that is used for
// SCR simnet bip32 public extended keys.
// Ecnodes to xsub in base58.
var SCRSimnetPublic = [4]byte{
	0x03,
	0x90,
	0x46,
	0x7d,
}

func toPublicVersion(version [4]byte) ([4]byte, error) {
	switch version {
	case BitcoinMainnetPrivate:
		return BitcoinMainnetPublic, nil
	case SCRMainnetPrivate:
		return SCRMainnetPublic, nil
	case SCRTestnetPrivate:
		return SCRTestnetPublic, nil
	case SCRpadevnetPrivate:
		return SCRpadevnetPublic, nil
	case SCRSimnetPrivate:
		return SCRSimnetPublic, nil
	}

	return [4]byte{}, errors.Errorf("unknown version %x", version)
}

func isPrivateVersion(version [4]byte) bool {
	switch version {
	case BitcoinMainnetPrivate:
		return true
	case SCRMainnetPrivate:
		return true
	case SCRTestnetPrivate:
		return true
	case SCRpadevnetPrivate:
		return true
	case SCRSimnetPrivate:
		return true
	}

	return false
}
