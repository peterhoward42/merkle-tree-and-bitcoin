package merkle

import (
	"github.com/peterhoward42/merkle-tree-and-bitcoin/src/hash"
	"math"
)

type MerkleTree struct {
	rows []Row
}

type MerklePath []hash.Byte32

type Row []hash.Byte32

func NewMerkleTree(bottomRow Row) (tree MerkleTree) {
	tree.rows = append(tree.rows, bottomRow)
	rowBeneath := bottomRow
	for {
		rowAbove := makeRowAbove(rowBeneath)
		tree.rows = append(tree.rows, rowAbove)
		rowBeneath = rowAbove
		if tree.isComplete() {
			break
		}
	}
	return
}

func (tree MerkleTree) MerkleRoot() hash.Byte32 {
	return tree.topRow()[0]
}

func (tree MerkleTree) MerklePathForLeaf(leafIndex int) (merklePath MerklePath) {
	i := leafIndex
	for _, row := range tree.rows[:len(tree.rows)-1] {
        siblingIndex := tree.siblingIndex(row, i)
		merklePath = append(merklePath, row[siblingIndex])
		i = i / 2 // Deliberate integer division
	}
	return
}

func CalculateMerkleRootFromMerklePath(
	leafHash hash.Byte32, merklePath MerklePath) hash.Byte32 {

	cumulativeHash := leafHash
	for _, hashInPath := range merklePath {
		cumulativeHash = hash.JoinAndHash(cumulativeHash, hashInPath)
	}
	return cumulativeHash
}

//---------------------------------------------------------------------------
// Private methods
//---------------------------------------------------------------------------

func (tree MerkleTree) isComplete() bool {
	return len(tree.topRow()) == 1
}

func (tree MerkleTree) topRow() Row {
	return tree.rows[len(tree.rows)-1]
}

// siblingIndex works out for a given node, which node in the same row should
// be considered its sibling. The sibling whose hash should be concatenated
// with its own hash that is, to create the parent node. Note however that 
// the nodes at the right hand end of rows with an odd number of elements do 
// not have one. In this special case, Merkle Trees, by defintion, substitute
// the hash of the first node for this role.
func (tree MerkleTree) siblingIndex(
        row Row, index int) (sibling int) {
	// For all odd indices, go left
	if (index % 2) == 1 {
		return index-1
	}
	// For most even indices, go right
	if (index + 1) <= len(row)-1 {
		return index+1
	}
	// Special case (required by definition of Merkle Tree)
	return index
}

func makeRowAbove(below Row) Row {
	size := int(math.Ceil(float64(len(below)) / 2.0))
	row := make([]hash.Byte32, size)
	for i, _ := range row {
		leftChild := i * 2
		rightChild := leftChild + 1
		if rightChild <= len(below)-1 {
			row[i] = hash.JoinAndHash(below[leftChild], below[rightChild])
		} else {
			row[i] = hash.JoinAndHash(below[leftChild], below[leftChild])
		}
	}
	return row
}
