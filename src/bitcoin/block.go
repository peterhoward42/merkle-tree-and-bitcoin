package bitcoin

import (
	"github.com/peterhoward42/merkle-tree-and-bitcoin/src/hash"
)

type Block struct {
	Records []Record
}

type BlockHeader struct {
	MerkleRoot hash.Byte32
}

type Record []byte

func (block Block) MakeListOfHashesForListOfRecords() (hashList []hash.Byte32) {
	for _, record := range block.Records {
		hashList = append(hashList, hash.Hash(record))
	}
	return
}
