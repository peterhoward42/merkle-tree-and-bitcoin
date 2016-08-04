package bitcoin

import (
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
	blockOfInterest int, recordToFetch int) {

	//blockHeader := spvNode.remote.GetBlockHeader(blockOfInterest)

	record, merklePath := spvNode.remote.GetRecord(
		blockOfInterest, recordToFetch)

	// Reproduce the merkle root calculation independently this end, using
	// the alleged record and the alleged Merkle Path.

	leafHash := hash.Hash(record)
	independentMerkleRoot := merkle.CalculateMerkleRootFromMerklePath(
		leafHash, merklePath)

	fmt.Printf("Indie root:\n%0x", independentMerkleRoot)
}
