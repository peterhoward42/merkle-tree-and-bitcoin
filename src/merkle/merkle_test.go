package merkle

import (
	"fmt"
	"github.com/peterhoward42/merkle-tree-and-bitcoin/src/hash"
	"testing"
)

func TestVerySmallTreeThatCanBeTracedByHand(t *testing.T) {
	// Build a tree that has bottom row comprising the hashes of just
	// three records, namely the ascii values that make up "abc".
	bottomRow := []hash.Byte32{}
	bottomRow = append(bottomRow, hash.Hash([]byte("a")))
	bottomRow = append(bottomRow, hash.Hash([]byte("b")))
	bottomRow = append(bottomRow, hash.Hash([]byte("c")))

	tree := NewMerkleTree(bottomRow)

	// Validate the shape of the tree.

	if len(tree.rows) != 3 {
		t.Errorf("Should be 3 rows.")
	}
	if len(tree.rows[0]) != 3 {
		t.Errorf("Row 0 should have 3 members")
	}
	if len(tree.rows[1]) != 2 {
		t.Errorf("Row 1 should have 3 members")
	}
	if len(tree.rows[2]) != 1 {
		t.Errorf("Row 2 should have 1 member")
	}

	// Validate the hashes in the middle row

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

	// Validate the hash in the top row, and that it has been captured as the
	// Merkle Root

	found = fmt.Sprintf("%0x", tree.rows[2][0])
	expected = fmt.Sprintf("%0x",
		hash.JoinAndHash(tree.rows[1][0], tree.rows[1][1]))
	if found != expected {
		t.Errorf(
			"Found:\n%s\ndiffers from expected:\n%s", found, expected)
	}
}
