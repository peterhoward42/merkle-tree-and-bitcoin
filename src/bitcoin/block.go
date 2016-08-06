package bitcoin

import (
	"github.com/peterhoward42/merkle-tree-and-bitcoin/src/hash"
)

// The Block type simulates a Bitcoin block - stripped down to only the parts
// we are interesed in for our program - which is just the records within it
// which would be transactions in the real Blockchain.
type Block struct {
	Records []Record
}

// The BlockHeader type simulates the Bitcoin block header - stripped down to
// only the parts we are interested in for our program - which is just the
// Merkle Root hash value.
type BlockHeader struct {
	MerkleRoot hash.Byte32
}

// The Record type simulates the byte-packed data used in the Blockchain to
// model transactions. For our purposes this need only be a chunk of contiguous
// bytes.
type Record []byte

// The MakeListOfHashesForListOfRecords() method provides a list of hash values
// that corresponds to the records held in the block.
func (block Block) MakeListOfHashesForListOfRecords() (hashList []hash.Byte32) {
	for _, record := range block.Records {
		hashList = append(hashList, hash.Hash(record))
	}
	return
}
