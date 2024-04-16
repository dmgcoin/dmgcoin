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
// Ecnodes to dmpr in base58.
var DmgcoinMainnetPrivate = [4]byte{
	0x64,
	0x6D,
	0x70,
	0x72,
}

// DmgcoinMainnetPublic is the version that is used for
// Dmgcoin mainnet bip32 public extended keys.
// Ecnodes to dmpb in base58.
var DmgcoinMainnetPublic = [4]byte{
	0x64,
	0x6D,
	0x70,
	0x62,
}

// DmgcoinTestnetPrivate is the version that is used for
// Dmgcoin testnet bip32 public extended keys.
// Ecnodes to dtpr in base58.
var DmgcoinTestnetPrivate = [4]byte{
	0x64,
	0x74,
	0x70,
	0x72,
}

// DmgcoinTestnetPublic is the version that is used for
// Dmgcoin testnet bip32 public extended keys.
// Ecnodes to dtpb in base58.
var DmgcoinTestnetPublic = [4]byte{
	0x64,
	0x74,
	0x70,
	0x62,
}

// DmgcoinDevnetPrivate is the version that is used for
// Dmgcoin devnet bip32 public extended keys.
// Ecnodes to ddpr in base58.
var DmgcoinDevnetPrivate = [4]byte{
	0x64,
	0x64,
	0x70,
	0x72,
}

// DmgcoinDevnetPublic is the version that is used for
// Dmgcoin devnet bip32 public extended keys.
// Ecnodes to ddpb in base58.
var DmgcoinDevnetPublic = [4]byte{
	0x64,
	0x64,
	0x70,
	0x62,
}

// DmgcoinSimnetPrivate is the version that is used for
// Dmgcoin simnet bip32 public extended keys.
// Ecnodes to dspr in base58.
var DmgcoinSimnetPrivate = [4]byte{
	0x64,
	0x73,
	0x70,
	0x72,
}

// DmgcoinSimnetPublic is the version that is used for
// Dmgcoin simnet bip32 public extended keys.
// Ecnodes to dspb in base58.
var DmgcoinSimnetPublic = [4]byte{
	0x64,
	0x73,
	0x70,
	0x62,
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
