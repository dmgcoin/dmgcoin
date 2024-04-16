// Copyright (c) 2014-2016 The btcsuite developers
// Use of this source code is governed by an ISC
// license that can be found in the LICENSE file.

package dagconfig

import (
	"math/big"

	"github.com/dmgcoin/dmgcoin/domain/consensus/model/externalapi"
	"github.com/dmgcoin/dmgcoin/domain/consensus/utils/blockheader"
	"github.com/dmgcoin/dmgcoin/domain/consensus/utils/subnetworks"
	"github.com/dmgcoin/dmgcoin/domain/consensus/utils/transactionhelper"
	"github.com/dmgcoin/go-muhash"
)

var genesisTxOuts = []*externalapi.DomainTransactionOutput{}

var genesisTxPayload = []byte{
	0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, // Blue score
	0x00, 0xE1, 0xF5, 0x05, 0x00, 0x00, 0x00, 0x00, // Subsidy
	0x00, 0x00, // Script version
	0x22,                                           // Varint
	0x00,                                           // OP_FALSE
	0x4D, 0x61, 0x67, 0x69, 0x63, 0x20, 0x69, 0x73, // Magic is science undiscovered
	0x20, 0x73, 0x63, 0x69, 0x65, 0x6E, 0x63, 0x65,
	0x20, 0x75, 0x6E, 0x64, 0x69, 0x73, 0x63, 0x6F,
	0x76, 0x65, 0x72, 0x65, 0x64,
}

// genesisCoinbaseTx is the coinbase transaction for the genesis blocks for
// the main network.
var genesisCoinbaseTx = transactionhelper.NewSubnetworkTransaction(0, []*externalapi.DomainTransactionInput{}, genesisTxOuts,
	&subnetworks.SubnetworkIDCoinbase, 0, genesisTxPayload)

// genesisHash is the hash of the first block in the block DAG for the main
// network (genesis block).
var genesisHash = externalapi.NewDomainHashFromByteArray(&[externalapi.DomainHashSize]byte{
	0x46, 0xe4, 0x4f, 0x81, 0x4, 0x38, 0xbd, 0x89,
	0x7, 0xa7, 0xb6, 0x8f, 0xe7, 0x1e, 0x71, 0xb8,
	0x7b, 0x7, 0x16, 0x8, 0x90, 0x52, 0xb0, 0x1f,
	0x1, 0xf8, 0xa5, 0x6e, 0x4f, 0x66, 0xc5, 0xb9,
})

// genesisMerkleRoot is the hash of the first transaction in the genesis block
// for the main network.
var genesisMerkleRoot = externalapi.NewDomainHashFromByteArray(&[externalapi.DomainHashSize]byte{
	0x16, 0x39, 0xc5, 0xcb, 0x4f, 0xe8, 0xf2, 0x8e,
	0x66, 0xcb, 0xc9, 0x7b, 0x7f, 0x13, 0x15, 0x40,
	0x8d, 0xfd, 0xc3, 0x66, 0x8, 0x46, 0xf2, 0x65,
	0x7b, 0xa7, 0x1e, 0x1b, 0x68, 0x8b, 0xd5, 0xb9,
})

// genesisBlock defines the genesis block of the block DAG which serves as the
// public transaction ledger for the main network.
var genesisBlock = externalapi.DomainBlock{
	Header: blockheader.NewImmutableBlockHeader(
		0,
		[]externalapi.BlockLevelParents{},
		genesisMerkleRoot,
		&externalapi.DomainHash{},
		externalapi.NewDomainHashFromByteArray(&[externalapi.DomainHashSize]byte{
			0x71, 0x0f, 0x27, 0xdf, 0x42, 0x3e, 0x63, 0xaa, 0x6c, 0xdb, 0x72, 0xb8, 0x9e, 0xa5, 0xa0, 0x6c, 0xff, 0xa3, 0x99, 0xd6, 0x6f, 0x16, 0x77, 0x04, 0x45, 0x5b, 0x5a, 0xf5, 0x9d, 0xef, 0x8e, 0x20,
		}),
		1637609671037,
		486722099,
		0x3392c,
		1312860, // Checkpoint DAA score
		0,
		big.NewInt(0),
		&externalapi.DomainHash{},
	),
	Transactions: []*externalapi.DomainTransaction{genesisCoinbaseTx},
}

var devnetGenesisTxOuts = []*externalapi.DomainTransactionOutput{}

var devnetGenesisTxPayload = []byte{
	0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, // Blue score
	0x00, 0xE1, 0xF5, 0x05, 0x00, 0x00, 0x00, 0x00, // Subsidy
	0x00, 0x00, // Script version
	0x22,                                           // Varint
	0x00,                                           // OP_FALSE
	0x4D, 0x61, 0x67, 0x69, 0x63, 0x20, 0x69, 0x73, // Magic is science undiscovered
	0x20, 0x73, 0x63, 0x69, 0x65, 0x6E, 0x63, 0x65,
	0x20, 0x75, 0x6E, 0x64, 0x69, 0x73, 0x63, 0x6F,
	0x76, 0x65, 0x72, 0x65, 0x64,
}

// devnetGenesisCoinbaseTx is the coinbase transaction for the genesis blocks for
// the development network.
var devnetGenesisCoinbaseTx = transactionhelper.NewSubnetworkTransaction(0,
	[]*externalapi.DomainTransactionInput{}, devnetGenesisTxOuts,
	&subnetworks.SubnetworkIDCoinbase, 0, devnetGenesisTxPayload)

// devGenesisHash is the hash of the first block in the block DAG for the development
// network (genesis block).
var devnetGenesisHash = externalapi.NewDomainHashFromByteArray(&[externalapi.DomainHashSize]byte{
	0x43, 0x65, 0x35, 0xe0, 0x82, 0x13, 0x80, 0xf0,
	0xa7, 0xe3, 0x4e, 0x4a, 0xd8, 0x1b, 0xed, 0x98,
	0xbd, 0x35, 0x43, 0x67, 0x84, 0x24, 0x8f, 0xe2,
	0xde, 0x33, 0x18, 0xae, 0xfd, 0xc3, 0x14, 0x4e,
})

// devnetGenesisMerkleRoot is the hash of the first transaction in the genesis block
// for the devopment network.
var devnetGenesisMerkleRoot = externalapi.NewDomainHashFromByteArray(&[externalapi.DomainHashSize]byte{
	0x16, 0x39, 0xc5, 0xcb, 0x4f, 0xe8, 0xf2, 0x8e,
	0x66, 0xcb, 0xc9, 0x7b, 0x7f, 0x13, 0x15, 0x40,
	0x8d, 0xfd, 0xc3, 0x66, 0x8, 0x46, 0xf2, 0x65,
	0x7b, 0xa7, 0x1e, 0x1b, 0x68, 0x8b, 0xd5, 0xb9,
})

// devnetGenesisBlock defines the genesis block of the block DAG which serves as the
// public transaction ledger for the development network.
var devnetGenesisBlock = externalapi.DomainBlock{
	Header: blockheader.NewImmutableBlockHeader(
		0,
		[]externalapi.BlockLevelParents{},
		devnetGenesisMerkleRoot,
		&externalapi.DomainHash{},
		externalapi.NewDomainHashFromByteArray(muhash.EmptyMuHashHash.AsArray()),
		0x11e9db49828,
		525264379,
		0x48e5e,
		0,
		0,
		big.NewInt(0),
		&externalapi.DomainHash{},
	),
	Transactions: []*externalapi.DomainTransaction{devnetGenesisCoinbaseTx},
}

var simnetGenesisTxOuts = []*externalapi.DomainTransactionOutput{}

var simnetGenesisTxPayload = []byte{
	0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, // Blue score
	0x00, 0xE1, 0xF5, 0x05, 0x00, 0x00, 0x00, 0x00, // Subsidy
	0x00, 0x00, // Script version
	0x01,                                                                   // Varint
	0x00,                                                                   // OP-FALSE
	0x6b, 0x61, 0x73, 0x70, 0x61, 0x2d, 0x73, 0x69, 0x6d, 0x6e, 0x65, 0x74, // dmgcoin-simnet
}

// simnetGenesisCoinbaseTx is the coinbase transaction for the simnet genesis block.
var simnetGenesisCoinbaseTx = transactionhelper.NewSubnetworkTransaction(0,
	[]*externalapi.DomainTransactionInput{}, simnetGenesisTxOuts,
	&subnetworks.SubnetworkIDCoinbase, 0, simnetGenesisTxPayload)

// simnetGenesisHash is the hash of the first block in the block DAG for
// the simnet (genesis block).
var simnetGenesisHash = externalapi.NewDomainHashFromByteArray(&[externalapi.DomainHashSize]byte{
	0x41, 0x1f, 0x8c, 0xd2, 0x6f, 0x3d, 0x41, 0xae,
	0xa3, 0x9e, 0x78, 0x57, 0x39, 0x27, 0xda, 0x24,
	0xd2, 0x39, 0x95, 0x70, 0x5b, 0x57, 0x9f, 0x30,
	0x95, 0x9b, 0x91, 0x27, 0xe9, 0x6b, 0x79, 0xe3,
})

// simnetGenesisMerkleRoot is the hash of the first transaction in the genesis block
// for the devopment network.
var simnetGenesisMerkleRoot = externalapi.NewDomainHashFromByteArray(&[externalapi.DomainHashSize]byte{
	0x19, 0x46, 0xd6, 0x29, 0xf7, 0xe9, 0x22, 0xa7,
	0xbc, 0xed, 0x59, 0x19, 0x05, 0x21, 0xc3, 0x77,
	0x1f, 0x73, 0xd3, 0x52, 0xdd, 0xbb, 0xb6, 0x86,
	0x56, 0x4a, 0xd7, 0xfd, 0x56, 0x85, 0x7c, 0x1b,
})

// simnetGenesisBlock defines the genesis block of the block DAG which serves as the
// public transaction ledger for the development network.
var simnetGenesisBlock = externalapi.DomainBlock{
	Header: blockheader.NewImmutableBlockHeader(
		0,
		[]externalapi.BlockLevelParents{},
		simnetGenesisMerkleRoot,
		&externalapi.DomainHash{},
		externalapi.NewDomainHashFromByteArray(muhash.EmptyMuHashHash.AsArray()),
		0x17c5f62fbb6,
		0x207fffff,
		0x2,
		0,
		0,
		big.NewInt(0),
		&externalapi.DomainHash{},
	),
	Transactions: []*externalapi.DomainTransaction{simnetGenesisCoinbaseTx},
}

var testnetGenesisTxOuts = []*externalapi.DomainTransactionOutput{}

var testnetGenesisTxPayload = []byte{
	0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, // Blue score
	0x00, 0xE1, 0xF5, 0x05, 0x00, 0x00, 0x00, 0x00, // Subsidy
	0x00, 0x00, // Script version
	0x01,                                                                         // Varint
	0x00,                                                                         // OP-FALSE
	0x6b, 0x61, 0x73, 0x70, 0x61, 0x2d, 0x74, 0x65, 0x73, 0x74, 0x6e, 0x65, 0x74, // dmgcoin-testnet
}

// testnetGenesisCoinbaseTx is the coinbase transaction for the testnet genesis block.
var testnetGenesisCoinbaseTx = transactionhelper.NewSubnetworkTransaction(0,
	[]*externalapi.DomainTransactionInput{}, testnetGenesisTxOuts,
	&subnetworks.SubnetworkIDCoinbase, 0, testnetGenesisTxPayload)

// testnetGenesisHash is the hash of the first block in the block DAG for the test
// network (genesis block).
var testnetGenesisHash = externalapi.NewDomainHashFromByteArray(&[externalapi.DomainHashSize]byte{
	0x9c, 0x3c, 0xb1, 0x6d, 0x54, 0x48, 0x54, 0x5,
	0xb0, 0xe, 0x19, 0x4b, 0x7f, 0x81, 0x22, 0xe6,
	0x91, 0x63, 0x5e, 0x83, 0x9, 0x76, 0xa9, 0x73,
	0x28, 0x4d, 0x47, 0xe1, 0xc7, 0xa9, 0x5e, 0xd0,
})

// testnetGenesisMerkleRoot is the hash of the first transaction in the genesis block
// for testnet.
var testnetGenesisMerkleRoot = externalapi.NewDomainHashFromByteArray(&[externalapi.DomainHashSize]byte{
	0x17, 0x34, 0x14, 0x8, 0xa5, 0x72, 0x45, 0x56,
	0x50, 0x4d, 0xf4, 0xd6, 0xcf, 0x51, 0x5c, 0xbf,
	0xbb, 0x22, 0x4, 0x30, 0xdc, 0x45, 0x1c, 0x74,
	0x3c, 0x22, 0xd5, 0xe9, 0x11, 0x72, 0xc, 0x2a,
})

// testnetGenesisBlock defines the genesis block of the block DAG which serves as the
// public transaction ledger for testnet.
var testnetGenesisBlock = externalapi.DomainBlock{
	Header: blockheader.NewImmutableBlockHeader(
		0,
		[]externalapi.BlockLevelParents{},
		testnetGenesisMerkleRoot,
		&externalapi.DomainHash{},
		externalapi.NewDomainHashFromByteArray(muhash.EmptyMuHashHash.AsArray()),
		0x661e104b, //1713246283, //0x14d2bd72185
		525264379,  // 525264379,
		0x1a14e,    //0x1a14e,
		0,
		0,
		big.NewInt(0),
		&externalapi.DomainHash{},
	),
	Transactions: []*externalapi.DomainTransaction{testnetGenesisCoinbaseTx},
}
