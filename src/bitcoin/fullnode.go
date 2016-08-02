package bitcoin

import (
	"github.com/peterhoward42/merkle-tree-and-bitcoin/src/merkle"
)

type FullBitcoinNode struct {
	DisHonest  bool
	block      Block
	merkleTree merkle.MerkleTree
}

func NewFullBitcoinNode() (node FullBitcoinNode) {
	node.block = MakeDummyBlockFromSherlockHolmesText()
	node.merkleTree = merkle.NewMerkleTree(node.block.HashList())
	return
}

func (node FullBitcoinNode) GetRecord42() (
	record Record, merklePath [][32]byte) {

	record = node.block.Records[42]
	merklePath = node.merkleTree.MerklePathForLeaf(42)

	return
}
