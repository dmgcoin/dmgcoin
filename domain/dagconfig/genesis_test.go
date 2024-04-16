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
	// 	0x16, 0x39, 0xc5, 0xcb, 0x4f, 0xe8, 0xf2, 0x8e,
	// 	0x66, 0xcb, 0xc9, 0x7b, 0x7f, 0x13, 0x15, 0x40,
	// 	0x8d, 0xfd, 0xc3, 0x66, 0x8, 0x46, 0xf2, 0x65,
	// 	0x7b, 0xa7, 0x1e, 0x1b, 0x68, 0x8b, 0xd5, 0xb9,
	// })

	//t.Logf("genesisHash %v", genesisHash)

	if !TestnetParams.GenesisHash.Equal(hash) {
		t.Fatalf("TestDevnetGenesisBlock: Genesis block hash does "+
			"not appear valid - got %v, want %v", hash,
			TestnetParams.GenesisHash)
	}
}
