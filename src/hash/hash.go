package hash

import (
	"crypto/sha256"
)

type Byte32 [32]byte

func Hash(input []byte) Byte32 {
	return sha256.Sum256(input)
}

func JoinAndHash(left Byte32, right Byte32) Byte32 {
	combined := left[:]
	combined = append(combined, right[:]...)
	return Hash(combined)
}
