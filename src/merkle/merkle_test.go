package merkle

import (
	"fmt"
	"github.com/peterhoward42/merkle-tree-and-bitcoin/src/hash"
	"testing"
)

/*
A note on the scope of these tests.

The code that these tests cover, exists to explain and to educate.
That is why it treats error handling as a distraction, and omits it.
So these tests exist only to ensure that the logic I expressed in the code is
doing what I intended, and produces correct results.
*/

func TestVerySmallTreeThatCanBeTracedByHand(t *testing.T) {
	tree := makeTreeUsingCharsInStringAsRecords("abc")

	// Validate the shape of the tree.

	if len(tree.rows) != 3 {
		t.Errorf("Wong number of rows")
	}
	if len(tree.rows[0]) != 3 {
		t.Errorf("Row has wrong length")
	}
	if len(tree.rows[1]) != 2 {
		t.Errorf("Row has wrong length")
	}
	if len(tree.rows[2]) != 1 {
		t.Errorf("Row has wrong length")
	}

	// Validate the hashes in the middle row, have the expected relationships
	// with those in the bottom row.

	found := fmt.Sprintf("%0x", tree.rows[1][0])
	expected := fmt.Sprintf("%0x",
		hash.JoinAndHash(tree.rows[0][0], tree.rows[0][1]))
	if found != expected {
		t.Errorf(
			"Found:\n%s\ndiffers from expected:\n%s", found, expected)
	}

	found = fmt.Sprintf("%0x", tree.rows[1][1])
	expected = fmt.Sprintf("%0x",
		hash.JoinAndHash(tree.rows[0][2], tree.rows[0][2]))
	if found != expected {
		t.Errorf(
			"Found:\n%s\ndiffers from expected:\n%s", found, expected)
	}

	// Validate the hash in the top row with respect to the middle row

	found = fmt.Sprintf("%0x", tree.rows[2][0])
	expected = fmt.Sprintf("%0x",
		hash.JoinAndHash(tree.rows[1][0], tree.rows[1][1]))
	if found != expected {
		t.Errorf(
			"Found:\n%s\ndiffers from expected:\n%s", found, expected)
	}

	// Validate the MerkleRoot query function.

	merkleRootFromQuery := fmt.Sprintf("%0x", tree.MerkleRoot())
	if merkleRootFromQuery != expected {
		t.Errorf(
			"Found:\n%s\ndiffers from expected:\n%s", found, expected)
	}
}

/*
When the number of elements in a Merkle Tree's bottom row is a power of two,
they form perfect binary trees in which every element in every non-leaf row
has two children. Otherwise we end up with some nodes on the right edge with
only a left child. If the number of leaf elements is one more than a power of
2, this is so for all the non leaf nodes at the right edge. For one less than a
power of two, it only occurs once, in the lowest, non leaf row.

This brings in singularities in the topology indexing arithmetic and
behaviour, so we will test the row lengths produced in each case.
*/

func TestPowerOfTwoRowLengths(t *testing.T) {
	tree := makeTreeUsingCharsInStringAsRecords("12345678")

	if len(tree.rows) != 4 {
		t.Errorf("Wong number of rows")
	}
	if len(tree.rows[0]) != 8 {
		t.Errorf("Row has wrong length")
	}
	if len(tree.rows[1]) != 4 {
		t.Errorf("Row has wrong length")
	}
	if len(tree.rows[2]) != 2 {
		t.Errorf("Row has wrong length")
	}
	if len(tree.rows[3]) != 1 {
		t.Errorf("Row has wrong length")
	}
}

func TestOneMoreThanPowerOfTwoRowLengths(t *testing.T) {
	tree := makeTreeUsingCharsInStringAsRecords("123456789")

	if len(tree.rows) != 5 {
		t.Errorf("Wong number of rows")
	}
	if len(tree.rows[0]) != 9 {
		t.Errorf("Row has wrong length")
	}
	if len(tree.rows[1]) != 5 {
		t.Errorf("Row has wrong length")
	}
	if len(tree.rows[2]) != 3 {
		t.Errorf("Row has wrong length")
	}
	if len(tree.rows[3]) != 2 {
		t.Errorf("Row has wrong length")
	}
	if len(tree.rows[4]) != 1 {
		t.Errorf("Row has wrong length")
	}
}

func TestOneLessThanPowerOfTwoRowLengths(t *testing.T) {
	tree := makeTreeUsingCharsInStringAsRecords("1234567")

	if len(tree.rows) != 4 {
		t.Errorf("Wong number of rows")
	}
	if len(tree.rows[0]) != 7 {
		t.Errorf("Row has wrong length")
	}
	if len(tree.rows[1]) != 4 {
		t.Errorf("Row has wrong length")
	}
	if len(tree.rows[2]) != 2 {
		t.Errorf("Row has wrong length")
	}
	if len(tree.rows[3]) != 1 {
		t.Errorf("Row has wrong length")
	}
}

func makeTreeUsingCharsInStringAsRecords(inputString string) MerkleTree {
	bottomRow := []hash.Byte32{}
	for _, c := range []byte(inputString) {
		bottomRow = append(bottomRow, hash.Hash([]byte{c}))
	}
	return NewMerkleTree(bottomRow)
}
