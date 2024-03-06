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

// DmgcoinMainnetPrivate is the version that is used for
// Dmgcoin mainnet bip32 private extended keys.
// Ecnodes to xprv in base58.
var DmgcoinMainnetPrivate = [4]byte{
	0x03,
	0x8f,
	0x2e,
	0xf4,
}

// DmgcoinMainnetPublic is the version that is used for
// Dmgcoin mainnet bip32 public extended keys.
// Ecnodes to kpub in base58.
var DmgcoinMainnetPublic = [4]byte{
	0x03,
	0x8f,
	0x33,
	0x2e,
}

// DmgcoinTestnetPrivate is the version that is used for
// Dmgcoin testnet bip32 public extended keys.
// Ecnodes to ktrv in base58.
var DmgcoinTestnetPrivate = [4]byte{
	0x03,
	0x90,
	0x9e,
	0x07,
}

// DmgcoinTestnetPublic is the version that is used for
// Dmgcoin testnet bip32 public extended keys.
// Ecnodes to ktub in base58.
var DmgcoinTestnetPublic = [4]byte{
	0x03,
	0x90,
	0xa2,
	0x41,
}

// DmgcoinDevnetPrivate is the version that is used for
// Dmgcoin devnet bip32 public extended keys.
// Ecnodes to kdrv in base58.
var DmgcoinDevnetPrivate = [4]byte{
	0x03,
	0x8b,
	0x3d,
	0x80,
}

// DmgcoinDevnetPublic is the version that is used for
// Dmgcoin devnet bip32 public extended keys.
// Ecnodes to xdub in base58.
var DmgcoinDevnetPublic = [4]byte{
	0x03,
	0x8b,
	0x41,
	0xba,
}

// DmgcoinSimnetPrivate is the version that is used for
// Dmgcoin simnet bip32 public extended keys.
// Ecnodes to ksrv in base58.
var DmgcoinSimnetPrivate = [4]byte{
	0x03,
	0x90,
	0x42,
	0x42,
}

// DmgcoinSimnetPublic is the version that is used for
// Dmgcoin simnet bip32 public extended keys.
// Ecnodes to xsub in base58.
var DmgcoinSimnetPublic = [4]byte{
	0x03,
	0x90,
	0x46,
	0x7d,
}

func toPublicVersion(version [4]byte) ([4]byte, error) {
	switch version {
	case BitcoinMainnetPrivate:
		return BitcoinMainnetPublic, nil
	case DmgcoinMainnetPrivate:
		return DmgcoinMainnetPublic, nil
	case DmgcoinTestnetPrivate:
		return DmgcoinTestnetPublic, nil
	case DmgcoinDevnetPrivate:
		return DmgcoinDevnetPublic, nil
	case DmgcoinSimnetPrivate:
		return DmgcoinSimnetPublic, nil
	}

	return [4]byte{}, errors.Errorf("unknown version %x", version)
}

func isPrivateVersion(version [4]byte) bool {
	switch version {
	case BitcoinMainnetPrivate:
		return true
	case DmgcoinMainnetPrivate:
		return true
	case DmgcoinTestnetPrivate:
		return true
	case DmgcoinDevnetPrivate:
		return true
	case DmgcoinSimnetPrivate:
		return true
	}

	return false
}
