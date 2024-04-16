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

/*func TestDevnetGenesisBlock2(t *testing.T) {
	// Check hash of the block against expected hash.
	hash := consensushashing.BlockHash(MainnetParams.GenesisBlock)
	trHash := consensushashing.TransactionHash(MainnetParams.GenesisBlock.Transactions[0])

	t.Logf("BLOCKHASH %v", hash)
	t.Logf("TXHASH %v", trHash)

	var testnetGenesisHash = externalapi.NewDomainHashFromByteArray(&[externalapi.DomainHashSize]byte{
		0x58, 0xc2, 0xd4, 0x19, 0x9e, 0x21, 0xf9, 0x10,
		0xd1, 0x57, 0x1d, 0x11, 0x49, 0x69, 0xce, 0xce,
		0xf4, 0x8f, 0x9, 0xf9, 0x34, 0xd4, 0x2c, 0xcb,
		0x6a, 0x28, 0x1a, 0x15, 0x86, 0x8f, 0x29, 0x99,
	})

	t.Logf("TESTNETGENESIS %v", testnetGenesisHash)

	if !MainnetParams.GenesisHash.Equal(hash) {
		t.Fatalf("TestDevnetGenesisBlock: Genesis block hash does "+
			"not appear valid - got %v, want %v", hash,
			MainnetParams.GenesisHash)
	}
}*/
