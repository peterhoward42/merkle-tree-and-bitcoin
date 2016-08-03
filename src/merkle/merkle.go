package merkle

import (
	"github.com/peterhoward42/merkle-tree-and-bitcoin/src/hash"
	"math"
)

type Row [][32]byte

type MerkleTree struct {
	MerkleRoot [32]byte
	rows       []Row
}

func NewMerkleTree(bottomRow Row) (tree MerkleTree) {
	tree.rows = append(tree.rows, bottomRow)
	for under := bottomRow; 
        len(under) > 1; under = tree.rows[len(tree.rows)-1] {
		tree.rows = append(tree.rows, makeRowAbove(under))
	}
	tree.MerkleRoot = tree.rows[len(tree.rows)-1][0]
	return
}

func (tree MerkleTree) MerklePathForLeaf(leafIndex int) (merklePath [][32]byte) {
	i := leafIndex
	for _, row := range tree.rows[:len(tree.rows)-1] {
		merklePath = append(merklePath, tree.siblingHash(row, i))
		i = i / 2 // Deliberate integer division
	}
	return
}

func CalculateRootFromPath(
	leafHash [32]byte, merklePath [][32]byte) (merkleRoot [32]byte) {

	cumulativeHash := leafHash
	for _, hashInPath := range merklePath {
		cumulativeHash = hash.JoinAndHash(cumulativeHash, hashInPath)
	}
	merkleRoot = cumulativeHash
	return
}

//---------------------------------------------------------------------------
// Private methods
//---------------------------------------------------------------------------

func (tree MerkleTree) siblingHash(row Row, index int) (hash [32]byte) {
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
	row := make([][32]byte, size)
	for i, _ := range row {
		leftChild := i * 2
		rightChild := leftChild + 1
		if rightChild < len(below)-1 {
			row[i] = hash.JoinAndHash(below[leftChild], below[rightChild])
		} else {
			row[i] = hash.JoinAndHash(below[leftChild], below[leftChild])
		}
	}
	return row
}
