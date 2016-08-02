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
	previous := bottomRow
	for {
		if len(previous) == 1 {
			tree.MerkleRoot = previous[0]
			break
		}
		tree.rows = append(tree.rows, makeRowAbove(previous))
		previous = tree.rows[len(tree.rows)-1]
	}
	return
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

func (tree MerkleTree) MerklePathForLeaf(leafIndex int) (merklePath [][32]byte) {
	i := leafIndex
	for _, row := range tree.rows[:len(tree.rows)-1] {
		merklePath = append(merklePath, tree.siblingHash(row, i))
		i = i / 2 // Deliberate integer division
	}
	return
}

func (tree MerkleTree) siblingHash(row Row, index int) (hash [32]byte) {
	if ((index % 2) == 0) && (index+1) <= (len(row)-1) {
		return row[index+1]
	}
	return row[index]
}
