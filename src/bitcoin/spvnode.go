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

func (spvNode SpvBitcoinNode) GetAndValidateRecord42FromRemoteNode(
	trustedMerkleRoot [32]byte) {

	record, merklePath := spvNode.remote.GetRecord42()

	// Reproduce the merkle root calculation independently this end, using
	// the alleged record and the alleged merkle path.

	leafHash := hash.Hash(record)
	myMerkleRoot := merkle.CalculateRootFromPath(leafHash, merklePath)

	fmt.Printf("Trusted:\n%0x", trustedMerkleRoot)
	fmt.Printf("Indie:\n%0x\n", myMerkleRoot)
}
