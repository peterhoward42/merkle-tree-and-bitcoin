package bitcoin

import (
	"errors"
	"fmt"
	"github.com/peterhoward42/merkle-tree-and-bitcoin/src/hash"
	"github.com/peterhoward42/merkle-tree-and-bitcoin/src/merkle"
)

type SpvBitcoinNode struct {
	remote *FullBitcoinNode
}

func NewSpvBitcoinNode(fullNode *FullBitcoinNode) (node *SpvBitcoinNode) {
	node = &SpvBitcoinNode{remote: fullNode}
	return
}

func (spvNode SpvBitcoinNode) FetchAndValidateRecordFromRemote(
	blockOfInterest int, recordToFetch int) (*Record, error) {

	blockHeader := spvNode.remote.GetBlockHeader(blockOfInterest)

	fetchedRecord, merklePath := spvNode.remote.GetRecord(
		blockOfInterest, recordToFetch)

	// Reproduce the merkle root calculation independently this end, using
	// the alleged record and the alleged Merkle Path.

	leafHash := hash.Hash(fetchedRecord)
	independentMerkleRoot := merkle.CalculateMerkleRootFromMerklePath(
		leafHash, merklePath)

	if independentMerkleRoot == blockHeader.MerkleRoot {
		return &fetchedRecord, nil
	} else {
		return nil, errors.New(
			formatErrorMessage(independentMerkleRoot,
				blockHeader.MerkleRoot))
	}
}

func formatErrorMessage(independentMerkleRoot hash.Byte32,
	blockHeaderMerkleRoot hash.Byte32) string {
	return fmt.Sprintf(
		"Independently calculated Merkle Root differs from "+
			"the one in the block header:\n%s\n%s",
		independentMerkleRoot.Hex(), blockHeaderMerkleRoot.Hex())
}
