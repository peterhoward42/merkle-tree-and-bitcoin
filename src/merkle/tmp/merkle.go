package merkle

import (
    "math")

type Row[][32]byte

type MerkleTree struct {
}

func NewMerkleTree(bottomRow Row) (tree MerkleTree) {
    return 
}

func makeRowAbove(below Row) Row {
    size := int(math.Ceil(float64(len(below)) / 2.0))
    row := make([][32]byte, size)
    for i, _ := range(row) {
        left := i * 2
        right := left + 1
        if right < len(below) - 1 {
            row[i] = joinAndHash(below[left], below[right])
        } else {
            row[i] = joinAndHash(below[left], below[left])
        }
    }
    return row
}

func joinAndHash(left [32]byte, right [32]byte) (hash [32]byte) {
    combined := left[:]
    combined = append(combined, right[:]...)
    return 
}

