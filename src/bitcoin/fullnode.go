package bitcoin

import (
	"github.com/peterhoward42/merkle-tree-and-bitcoin/src/hash"
	"github.com/peterhoward42/merkle-tree-and-bitcoin/src/merkle"
)

type FullBitcoinNode struct {
	DisHonest  bool
	block      Block
	merkleTree merkle.MerkleTree
}

func NewFullBitcoinNode() (*FullBitcoinNode) {
    node := FullBitcoinNode{}
	node.block = MakeDummyBlockFromSherlockHolmesText()

    hashesForBottomRowOfTree := node.block.GetHashesForAllRecords()
	node.merkleTree = merkle.NewMerkleTree(hashesForBottomRowOfTree)
	return &node
}

func (node FullBitcoinNode) MerkleRootForBlock() (hash.Byte32) {
	return node.merkleTree.MerkleRoot()
}

func (node FullBitcoinNode) GetRecord42() (
	record Record, 
    merklePath []hash.Byte32) {

	record = node.block.Records[42]
	merklePath = node.merkleTree.MerklePathForLeaf(42)
	return
}
