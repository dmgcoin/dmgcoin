// Copyright (c) 2014-2016 The btcsuite developers
// Use of this source code is governed by an ISC
// license that can be found in the LICENSE file.

package dagconfig

import (
	"testing"

	"github.com/dmgcoin/dmgcoin/domain/consensus/utils/consensushashing"
)

// TestGenesisBlock tests the genesis block of the main network for validity by
// checking the encoded hash.
func TestGenesisBlock(t *testing.T) {
	// Check hash of the block against expected hash.
	hash := consensushashing.BlockHash(MainnetParams.GenesisBlock)
	if !MainnetParams.GenesisHash.Equal(hash) {
		t.Fatalf("TestGenesisBlock: Genesis block hash does "+
			"not appear valid - got %v, want %v", hash, MainnetParams.GenesisHash)
	}
}

// TestTestnetGenesisBlock tests the genesis block of the test network for
// validity by checking the hash.
func TestTestnetGenesisBlock(t *testing.T) {
	// Check hash of the block against expected hash.
	hash := consensushashing.BlockHash(TestnetParams.GenesisBlock)
	if !TestnetParams.GenesisHash.Equal(hash) {
		t.Fatalf("TestTestnetGenesisBlock: Genesis block hash does "+
			"not appear valid - got %v, want %v", hash,
			TestnetParams.GenesisHash)
	}
}

// TestSimnetGenesisBlock tests the genesis block of the simulation test network
// for validity by checking the hash.
func TestSimnetGenesisBlock(t *testing.T) {
	// Check hash of the block against expected hash.
	hash := consensushashing.BlockHash(SimnetParams.GenesisBlock)
	if !SimnetParams.GenesisHash.Equal(hash) {
		t.Fatalf("TestSimnetGenesisBlock: Genesis block hash does "+
			"not appear valid - got %v, want %v", hash,
			SimnetParams.GenesisHash)
	}
}

// TestDevnetGenesisBlock tests the genesis block of the development network
// for validity by checking the encoded hash.
func TestDevnetGenesisBlock(t *testing.T) {
	// Check hash of the block against expected hash.
	hash := consensushashing.BlockHash(DevnetParams.GenesisBlock)

	if !DevnetParams.GenesisHash.Equal(hash) {
		t.Fatalf("TestDevnetGenesisBlock: Genesis block hash does "+
			"not appear valid - got %v, want %v", hash,
			DevnetParams.GenesisHash)
	}
}

func TestDevnetGenesisBlock2(t *testing.T) {
	// Check hash of the block against expected hash.
	hash := consensushashing.BlockHash(TestnetParams.GenesisBlock)
	trHash := consensushashing.TransactionHash(TestnetParams.GenesisBlock.Transactions[0])

	t.Logf("BLOCKHASH %v", hash)
	t.Logf("TXHASH %v", trHash)

	// var genesisHash = externalapi.NewDomainHashFromByteArray(&[externalapi.DomainHashSize]byte{
	// 	0xe6, 0xe2, 0x1, 0x39, 0xf6, 0x2a, 0xfe, 0xc,
	// 	0x31, 0xf8, 0xcc, 0x9a, 0x1c, 0x14, 0x6e, 0x1,
	// 	0x1f, 0xb, 0xb7, 0xa5, 0xd2, 0x6, 0xa7, 0xba,
	// 	0x88, 0xe9, 0xe6, 0x5e, 0x7, 0xe1, 0x37, 0xe6,
	// })

	// t.Logf("genesisHash %v", genesisHash)

	if !TestnetParams.GenesisHash.Equal(hash) {
		t.Fatalf("TestDevnetGenesisBlock: Genesis block hash does "+
			"not appear valid - got %v, want %v", hash,
			TestnetParams.GenesisHash)
	}
}
