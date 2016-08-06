package bitcoin

import (
	"errors"
	"fmt"
	"github.com/peterhoward42/merkle-tree-and-bitcoin/src/hash"
	"github.com/peterhoward42/merkle-tree-and-bitcoin/src/merkle"
)

// SpvBitcoinNode is our representation of a Single Payment Verification (SPV)
// Bitcoin node. It holds a reference to a FullBitcoinNode that it will use as
// its peer, and can call methods on it to simulate the passing of network
// messages in the real Bitcoin protocol. It exposes a single API method itself,
// which clients can use to mandate it to fetch a record from its remote
// peer, and verify the record's integrity and authenticity using a Merkle
// Path.
type SpvBitcoinNode struct {
	Remote *FullBitcoinNode
}

// FetchAndValidateRecordFromRemote is an exposed API method that clients can
// use to instruct the SPV node to fetch a nominated record from its remote
// full node peer and to verify the record's integrity and authenticity using
// the Merkle Path returned by the remote node with the record. You specify the
// block and record you want using indices. If the Merkle Path-baesd
// verification fails, then the error returned object is non nil and will
// contain a description of the error.
func (spvNode SpvBitcoinNode) FetchAndValidateRecordFromRemote(
	blockOfInterest int, recordToFetch int) (*Record, error) {

	// The process is initialised by requesting the block header from the
	// remote node. A real Bitcoin client will check several things based on
	// the block header - which are explained in the accompanying article. But
	// for the purposes of this example code, we are only interested in the
	// fact it contains what the remote node claims to be the Merkle Root
	// hash value for the block. We will use this as part of our validation
	// shortly.
	blockHeader := spvNode.Remote.GetBlockHeader(blockOfInterest)

	// Next we ask the remote node to provide the record, along with the Merkle
	// Path for it.
	fetchedRecord, merklePath := spvNode.Remote.GetRecord(
		blockOfInterest, recordToFetch)

	// Now we can use the Merkle Path to prove to ourselves that the remote
	// node is telling the truth by starting with a hash for the returned
	// record that we calculate independently for ourselves, and then hashing
	// our way along the Merkle Path to make sure we end up with the same
	// Merkle Root value as the one captured in the block header before we
	// started.
	leafHash := hash.Hash(fetchedRecord)
	independentMerkleRoot := merkle.CalculateMerkleRootFromMerklePath(
		leafHash, merklePath)

	// And depending on our decision about if the remote node is telling the
	// truth, we either return the record or produce an error return value.
	if independentMerkleRoot == blockHeader.MerkleRoot {
		return &fetchedRecord, nil
	} else {
		return nil, errors.New(
			formatErrorMessage(independentMerkleRoot,
				blockHeader.MerkleRoot))
	}
}

// formatErrorMessage is a helper function used to remove distracting
// boiler-plate code from the FetchAndValidateRecord method.
func formatErrorMessage(independentMerkleRoot hash.Byte32,
	blockHeaderMerkleRoot hash.Byte32) string {
	return fmt.Sprintf(
		"Independently calculated Merkle Root differs from "+
			"the one in the block header:\n%s\n%s",
		independentMerkleRoot.Hex(), blockHeaderMerkleRoot.Hex())
}
