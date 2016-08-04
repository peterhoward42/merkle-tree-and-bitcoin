package merkle

import (
	"fmt"
	"github.com/peterhoward42/merkle-tree-and-bitcoin/src/hash"
	"testing"
)

/*
A note on the scope of these tests.

The code that these tests cover, exists to explain and to educate, not to be an
industrialised solution. That is why for example, error handling is omitted
in the interests of simplicity.

So the tests exist solely to check that the code is implementing the logic it
is intended to, and to provide debugging support during the development and 
future extension, by me or others.
*/

// TestConstructionOfVerySmallTreeThatCanBeTracedByHand looks in detail at the
// shape and contents of a Merkle Tree that is so small a human being can
// readily assimilate what should be there. 
func TestConstructionOfVerySmallTreeThatCanBeTracedByHand(t *testing.T) {
	tree := makeTreeUsingCharsInStringAsRecords("abc")

	// Verify the shape of the tree.

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

	// Verify the hashes in the middle row, have the expected relationships
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

	// Verify the hash in the top row (the Merkle Root) with respect to the 
    // middle row

	found = fmt.Sprintf("%0x", tree.rows[2][0])
	expected = fmt.Sprintf("%0x",
		hash.JoinAndHash(tree.rows[1][0], tree.rows[1][1]))
	if found != expected {
		t.Errorf(
			"Found:\n%s\ndiffers from expected:\n%s", found, expected)
	}

	// Verify the MerkleRoot query function.

	merkleRootFromQuery := fmt.Sprintf("%0x", tree.MerkleRoot())
	if merkleRootFromQuery != expected {
		t.Errorf(
			"Found:\n%s\ndiffers from expected:\n%s", found, expected)
	}
}

// TestPowerOfTwoRowLengths exercises the construction of a Merkle Tree when
// the number of elements in the bottom row is a power of two. This is
// significant because in this case the tree will be perfect binary tree, with
// all non-leaf nodes having both a left and right child. The checks are to
// make sure that there are the expected number of rows, each with the expected
// number of nodes.
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

// TestOneMoreThanPowerOfTwoRowLengths exercises the construction of a Merkle 
// Tree when the number of elements in the bottom row exceeds a power of two by
// one. This is significant because in this case the tree will have elements
// all down its right hand side that have only a left child, and this property
// stimulates different paths in both the row relationship arithmetic and
// the choice of nodes to combine for hashing.
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

// TestOneLess exercises the construction of a Merkle Tree when the number of 
// elements in the bottom row is less than a power of two by one. This is 
// significant because in this case the tree will have just one node in in the
// first row above the leaves that has only a left child.
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

// TestActualHashValuesInSmallestViableTree makes a Merkle Tree with just two
// elements in the bottom row comprising the hashes of the ASCII characters
// 'A' and 'B'. This allows us to verify the hash values created for both leaf
// nodes and the only node present - the root.
func TestActualHashValuesInSmallestViableTree(t *testing.T) {
    // SHA256 hashes reference values available here:
    // http://www.xorbin.com/tools/sha256-hash-calculator 
	tree := makeTreeUsingCharsInStringAsRecords("AB")

	leftLeafHash := fmt.Sprintf("%0x", tree.rows[0][0])
    expected :=
    "559aead08264d5795d3909718cdd05abd49572e84fe55590eef31a88a08fdffd"
	if leftLeafHash != expected {
		t.Errorf(
			"Found:\n%s\ndiffers from expected:\n%s", leftLeafHash, expected)
	}

	rightLeafHash := fmt.Sprintf("%0x", tree.rows[0][1])
    expected =
    "df7e70e5021544f4834bbee64a9e3789febc4be81470df629cad6ddb03320a5c"
	if rightLeafHash != expected {
		t.Errorf(
			"Found:\n%s\ndiffers from expected:\n%s", rightLeafHash, expected)
	}
}

// makeTreeUsingCharsInStringAsRecords is a utility function to support unit
// tests. It creates Merkle Trees in which the leaf nodes are the hashes of 
// single bytes that comprise a string.
func makeTreeUsingCharsInStringAsRecords(inputString string) MerkleTree {
	bottomRow := []hash.Byte32{}
	for _, c := range []byte(inputString) {
		bottomRow = append(bottomRow, hash.Hash([]byte{c}))
	}
	return NewMerkleTree(bottomRow)
}
