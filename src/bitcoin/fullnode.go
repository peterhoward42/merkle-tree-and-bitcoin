package bitcoin

import (
	"github.com/peterhoward42/merkle-tree-and-bitcoin/src/merkle"
)

// FullBitcoinNode is our representation of a real full Bitcoin node.
// It holds a mock Blockchain, along with corresponding block headers, and
// Merkle Trees. It exposes a small set of query methods to simulate the
// receipt of messages from a Single Payment Verification (SPV) Bitcoin node.
type FullBitcoinNode struct {
	// These slices are indexed in lock step with each other.
	blockChain   []Block
	blockHeaders []BlockHeader
	merkleTrees  []merkle.MerkleTree
}

// The NewFullBitcoinNode function is a factory that produces a
// FullBitcoinNode, including populating it with its mock Blockchain, and
// building the corresponding block headers and Merkle Trees.
func NewFullBitcoinNode() (node FullBitcoinNode) {
	node.populateBlockChainWithMadeUpBlocks()
	node.populateMerkleTrees()
	node.populateBlockHeaders()
	return node
}

// GetBlockHeader is a query function to retreive a block header by its index
// position in the Blockchain.
func (node FullBitcoinNode) GetBlockHeader(blockOfInterest int) BlockHeader {
	return node.blockHeaders[blockOfInterest]
}

// GetRecord is a query function to retreive a record (transaction) specified
// by the index of a block, and the index of the record required from it.
// In addition to the record, it returns also the Merkle Path that can be used
// by the caller to verify both the block's internal integrity but also its
// authenticity as the record that lives in that slot in the rest of the block.
func (node FullBitcoinNode) GetRecord(
	blockOfInterest int,
	recordToFetch int) (record Record, merklePath merkle.MerklePath) {

	block := node.blockChain[blockOfInterest]
	record = block.Records[recordToFetch]

	// We assemble the Merkle Path on-the-fly by traversing the Merkle Tree
	// held for this block.
	merkleTree := node.merkleTrees[blockOfInterest]
	merklePath = merkleTree.MerklePathForLeaf(recordToFetch)
	return
}

//---------------------------------------------------------------------------
// Private below
//---------------------------------------------------------------------------

// populateBlockChainWithMadeUpBlocks is a method that builds some blocks
// to put into the Blockchain based on some text contents it downloads
// from Sherlock Holmes novels.
func (node *FullBitcoinNode) populateBlockChainWithMadeUpBlocks() {
	blocks := MakeSetOfBlocksBasedOnContentsOfDownloadedBooks()
	node.blockChain = append(node.blockChain, blocks...)
}

// populateMerkleTrees is a method that builds and installs a Merkle Tree for
// each of the blocks present in the Blockchain.
func (node *FullBitcoinNode) populateMerkleTrees() {
	for _, block := range node.blockChain {
		bottomRow := block.MakeListOfHashesForListOfRecords()
		tree := merkle.NewMerkleTree(bottomRow)
		node.merkleTrees = append(node.merkleTrees, tree)
	}
}

// populateBlockHeaders is a method that builds and installs a Block Header
// for each block in the Blockchain.
func (node *FullBitcoinNode) populateBlockHeaders() {
	for _, merkleTree := range node.merkleTrees {
		merkleRoot := merkleTree.MerkleRoot()
		blockHeader := BlockHeader{MerkleRoot: merkleRoot}
		node.blockHeaders = append(node.blockHeaders, blockHeader)
	}
}
