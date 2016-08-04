package bitcoin

import (
	"github.com/peterhoward42/merkle-tree-and-bitcoin/src/hash"
)

type Record []byte

type Block struct {
	Records []Record
}

func (block Block) GetHashesForAllRecords() (hashList []hash.Byte32) {
	return
}
