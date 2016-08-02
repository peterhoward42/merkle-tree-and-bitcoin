package hash

import (
	"crypto/sha256"
)

func Hash(input []byte) [32]byte {
	return sha256.Sum256(input)
}

func JoinAndHash(left [32]byte, right [32]byte) [32]byte {
	combined := left[:]
	combined = append(combined, right[:]...)
	return Hash(combined)
}
