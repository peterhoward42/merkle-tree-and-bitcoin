package bitcoin

import (
	"github.com/peterhoward42/merkle-tree-and-bitcoin/src/merkle"
	"github.com/peterhoward42/merkle-tree-and-bitcoin/src/hash"
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

// Assumed to be trusted
func (node FullBitcoinNode) MerkleRootForBlock() (merkleRoot hash.Byte32) {
	return node.merkleTree.MerkleRoot()
}

func (node FullBitcoinNode) GetRecord42() (
	record Record, merklePath []hash.Byte32) {

	record = node.block.Records[42]
	merklePath = node.merkleTree.MerklePathForLeaf(42)
	return
}
