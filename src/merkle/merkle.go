package merkle

import (
	"github.com/peterhoward42/merkle-tree-and-bitcoin/src/hash"
	"math"
)

type MerkleTree struct {
	rows []Row
}

type Row []hash.Byte32

type MerklePath []MerklePathElement

type MerklePathElement struct {
	hash                    hash.Byte32
	useFirstInConcatenation bool // me+other, not other+me
}

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
	merklePath MerklePath) {
	i := leafIndex
	for _, row := range tree.rows[:len(tree.rows)-1] {
		sibling, useFirstInConcatenation := row.evaluateSibling(i)
		merklePathElement := MerklePathElement{
			hash: row[sibling],
			useFirstInConcatenation: useFirstInConcatenation}
		merklePath = append(merklePath, merklePathElement)
		i = i / 2 // Deliberate integer division
	}
	return
}

func CalculateMerkleRootFromMerklePath(
	leafHash hash.Byte32, merklePath MerklePath) hash.Byte32 {

	cumulativeHash := leafHash
	for _, merklePathElement := range merklePath {
		if merklePathElement.useFirstInConcatenation {
			cumulativeHash = hash.JoinAndHash(
				merklePathElement.hash, cumulativeHash)
		} else {
			cumulativeHash = hash.JoinAndHash(
				cumulativeHash, merklePathElement.hash)
		}
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

/* evaluateSibling works out which neighbour of element X in a table row is
 * the sibling whose hash should be combined with that of X to form the parent
 * node hash. There is a general rule depending on if X is an even or odd
 * numbered element, plus a special case for X being at the right hand end of
 * an odd length row. In addition to identifying the sibling, this function
 * must capture the sequence in which the hash concatenations must be done, and
 * hence returns both the sibling index and a flag to signify the concatenation
 * order required. */
func (row Row) evaluateSibling(myIndex int) (
	siblingIndex int, useFirstInConcatenation bool) {

	// For all odd indices, the pair is leftNeighbour->me.
	// For most even indices, the pair is me->rightNeighbour
	// For special case, the pair is me->me (by definition in Merkle Trees)

	if myIndex%2 == 1 {
		siblingIndex = myIndex - 1
		useFirstInConcatenation = true
	} else if (myIndex + 1) <= len(row)-1 {
		siblingIndex = myIndex + 1
		useFirstInConcatenation = false
	} else {
		siblingIndex = myIndex
		useFirstInConcatenation = true // moot
	}
	return
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
