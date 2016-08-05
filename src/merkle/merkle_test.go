package merkle

import (
	"github.com/peterhoward42/merkle-tree-and-bitcoin/src/hash"
	"testing"
)

/*
A note on the scope of these tests.

The code that these tests cover, exists to explain and to educate, not to be an
industrialised solution. That is why for example, error handling is omitted
to make the code easier to understand.

So the tests exist solely to check that the code is implementing the logic it
is intended to, and to provide debugging support for developers.
*/

// These constants are reference SHA-256 hash values, sourced from
// http://www.xorbin.com/tools/sha256-hash-calculator
const hashOfLetterA = "559aead08264d5795d3909718cdd05abd49572e84fe55590eef31a88a08fdffd"
const hashOfLetterB = "df7e70e5021544f4834bbee64a9e3789febc4be81470df629cad6ddb03320a5c"
const hashOfLetterC = "6b23c0d5f35d1b11f9b683f0b0a617355deb11277d91ae091d399c655b87940d"

// TestShapeOfABCTree builds a tree in which the leaves are the hash values
// of the characters A B C respectively, and then ensures the number and length
// of the rows in the tree are as expected.
func TestShapeOfABCTree(t *testing.T) {
	tree := makeABCTree()

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
}

// TestLeafHashesInABCTree builds a tree in which the leaves are the hash values
// of the characters A B C respectively, and then ensures the leaf hash values
// are the reference values expected.
func TestLeafHashesInABCTree(t *testing.T) {
	tree := makeABCTree()

	found := tree.rows[0][0].Hex()
	expected := hashOfLetterA
	if found != expected {
		t.Errorf(
			"Found:\n%s\ndiffers from expected:\n%s", found, expected)
	}

	found = tree.rows[0][1].Hex()
	expected = hashOfLetterB
	if found != expected {
		t.Errorf(
			"Found:\n%s\ndiffers from expected:\n%s", found, expected)
	}

	found = tree.rows[0][2].Hex()
	expected = hashOfLetterC
	if found != expected {
		t.Errorf(
			"Found:\n%s\ndiffers from expected:\n%s", found, expected)
	}
}

// TestSiblingIdentification exercises the function inside the Merkle Tree
// implementation that decides which node should be used to provide the hash
// value for the node's sibling. This has 3 separate logical paths, each of
// which are stimulated and checked.
func TestSiblingIdentification(t *testing.T) {
	tree := makeABCTree()

    // Odd numbered elements should use their left neighbour.
    siblingIndex := tree.siblingIndex(tree.rows[1], 1)
    if siblingIndex != 0 {
		t.Errorf("Wrong sibling index")
    }
}

// TestHashRelationshipsInsideABCTree builds a tree in which the leaves are
// the hash values of the characters A B C respectively, and then ensures the
// hash relationships between parents and children are as expected.
func TestHashRelationshipsInsideABCTree(t *testing.T) {
	tree := makeABCTree()

	// Middle row WRT bottom row.
	found := tree.rows[1][0].Hex()
	expected := hash.JoinAndHash(tree.rows[0][0], tree.rows[0][1]).Hex()
	if found != expected {
		t.Errorf(
			"Found:\n%s\ndiffers from expected:\n%s", found, expected)
	}

	found = tree.rows[1][1].Hex()
	expected = hash.JoinAndHash(tree.rows[0][2], tree.rows[0][2]).Hex()
	if found != expected {
		t.Errorf(
			"Found:\n%s\ndiffers from expected:\n%s", found, expected)
	}

	// Top row WRT middle row.
	found = tree.rows[2][0].Hex()
	expected = hash.JoinAndHash(tree.rows[1][0], tree.rows[1][1]).Hex()
	if found != expected {
		t.Errorf(
			"Found:\n%s\ndiffers from expected:\n%s", found, expected)
	}
}

// TestMerkleRootQueryForABCTree builds a tree in which the leaves are
// the hash values of the characters A B C respectively, and then ensures the
// MerkleRoot query function yields the top node value.
func TestMerkleRootQueryForABCTree(t *testing.T) {
	tree := makeABCTree()
	topNode := tree.rows[2][0].Hex()
	queriedRoot := tree.MerkleRoot().Hex()
	if queriedRoot != topNode {
		t.Errorf(
			"Queried root:\n%s\ndiffers from top node:\n%s",
			queriedRoot, topNode)
	}
}

// TestPowerOfTwoRowLengths exercises the construction of a Merkle Tree when
// the number of elements in the bottom row is a power of two. This is
// significant because in this case the tree will be perfect binary tree, with
// all non-leaf nodes having both a left and right child. The checks are to
// make sure that there are the expected number of rows, each with the expected
// number of nodes.
func TestPowerOfTwoRowLengths(t *testing.T) {
	tree := makeTreeUsingEachCharInStringAsRecord("12345678")

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
	tree := makeTreeUsingEachCharInStringAsRecord("123456789")

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

	tree := makeTreeUsingEachCharInStringAsRecord("1234567")

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

// makeABCTree is a utility function that creates a tree in which the leaves
// are the hash values for the characters 'A', 'B', 'C' respectively.
func makeABCTree() MerkleTree {
	return makeTreeUsingEachCharInStringAsRecord("ABC")
}

// makeTreeUsingEachCharInStringAsRecord is a utility function to support unit
// tests. It creates Merkle Trees in which the leaf nodes are the hashes of
// single bytes taken from a string.
func makeTreeUsingEachCharInStringAsRecord(inputString string) MerkleTree {
	bottomRow := []hash.Byte32{}
	for _, c := range []byte(inputString) {
		bottomRow = append(bottomRow, hash.Hash([]byte{c}))
	}
	return NewMerkleTree(bottomRow)
}
