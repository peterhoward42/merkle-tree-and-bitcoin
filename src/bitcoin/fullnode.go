package bitcoin

import (
	"github.com/peterhoward42/merkle-tree-and-bitcoin/src/merkle"
)

type FullBitcoinNode struct {
	// Indexed in lock step
	blockChain   []Block
	blockHeaders []BlockHeader
	merkleTrees  []merkle.MerkleTree
}

func NewFullBitcoinNode() (node FullBitcoinNode) {
	node.populateBlockChainWithMadeUpBlocks()
	node.populateMerkleTrees()
	node.populateBlockHeaders()
	return node
}

func (node FullBitcoinNode) GetBlockHeader(blockOfInterest int) BlockHeader {
	return node.blockHeaders[blockOfInterest]
}

func (node FullBitcoinNode) GetRecord(
	blockOfInterest int,
	recordToFetch int) (record Record, merklePath merkle.MerklePath) {

	block := node.blockChain[blockOfInterest]
	record = block.Records[recordToFetch]
	merkleTree := node.merkleTrees[blockOfInterest]
	merklePath = merkleTree.MerklePathForLeaf(recordToFetch)
	return
}

//---------------------------------------------------------------------------
// Private below
//---------------------------------------------------------------------------

func (node *FullBitcoinNode) populateBlockChainWithMadeUpBlocks() {
	sherlockHolmesStories := []string{"bosc", "cree", "danc", "gold"}
	for _, urlFragment := range sherlockHolmesStories {
		block := MakeBlockBasedOnBookText(urlFragment)
		node.blockChain = append(node.blockChain, block)
	}
}

func (node *FullBitcoinNode) populateMerkleTrees() {
	for _, block := range node.blockChain {
		bottomRow := block.MakeListOfHashesForListOfRecords()
		tree := merkle.NewMerkleTree(bottomRow)
		node.merkleTrees = append(node.merkleTrees, tree)
	}
}

func (node *FullBitcoinNode) populateBlockHeaders() {
	for _, merkleTree := range node.merkleTrees {
		merkleRoot := merkleTree.MerkleRoot()
		blockHeader := BlockHeader{MerkleRoot: merkleRoot}
		node.blockHeaders = append(node.blockHeaders, blockHeader)
	}
}
