package bitcoin

import ()

type Record []byte

type Block struct {
	Records []Record
}

func (block Block) HashList() (hashList [][32]byte) {
	return
}
