package merkle

import (
	"github.com/peterhoward42/merkle-tree-and-bitcoin/src/hash"
	"math"
)

type MerkleTree struct {
	rows []Row
}

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

func (tree MerkleTree) MerklePathForLeaf(leafIndex int) (
	merklePath []hash.Byte32) {
	i := leafIndex
	for _, row := range tree.rows[:len(tree.rows)-1] {
		merklePath = append(merklePath, tree.siblingHash(row, i))
		i = i / 2 // Deliberate integer division
	}
	return
}

func CalculateMerkleRootFromMerklePath(
	leafHash hash.Byte32, merklePath []hash.Byte32) hash.Byte32 {

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

func (tree MerkleTree) siblingHash(row Row, index int) (hash hash.Byte32) {
	// For all odd indices, go left
	if (index % 2) == 1 {
		return row[index-1]
	}
	// For most even indices, go right
	if (index + 1) <= len(row)-1 {
		return row[index+1]
	}
	// Special case (required by definition of Merkle Tree)
	return row[index]
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
