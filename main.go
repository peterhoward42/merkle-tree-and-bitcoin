package main

import (
	"fmt"
	"net/http"
	"io/ioutil"
)

func main() {
	fullNode := NewFullNode()
	spvNode := NewSpvNode(fullNode)

	spvNode.GetAndValidateRecord42FromRemoteNode() // Prints an OK message
	fullNode.GoDishonest()
	spvNode.GetAndValidateRecord42FromRemoteNode() // Prints a failure message

	fmt.Printf("Finished\n")
}

//----------------------------------------------------------------------------------------------
// Full Node
//----------------------------------------------------------------------------------------------

type FullNode struct {
	disHonest bool
	Block Block
}

func NewFullNode() *FullNode {
	fullNode := &FullNode{}
	fullNode.Block.Build()
	return fullNode
}

func (fullNode *FullNode) GoDishonest() {
	fullNode.disHonest = true
}

//----------------------------------------------------------------------------------------------
// SPV Node
//----------------------------------------------------------------------------------------------

type SpvNode struct {
	remote *FullNode
}

func NewSpvNode(remote *FullNode) *SpvNode {
	spvNode := &SpvNode{remote: remote}
	return spvNode
}

func (spvNode *SpvNode) GetAndValidateRecord42FromRemoteNode() {
	record, merklePath := spvNode.remote.Block.GetRecord(42)
	fmt.Printf("Record: %v", record)
	fmt.Printf("Path: %v", merklePath)
}

//----------------------------------------------------------------------------------------------
// Block
//----------------------------------------------------------------------------------------------

type Block struct {
	records [][]byte
	// We use the customary index arithmetic to store our full binary tree in a flat
	// sequence. Where the left child of item N lives at N * 2 + 1, with the
	// corresponding right child at N * 2 + 2
	merkleTree []int
}



func (blk *Block) Build() {
	blk.fillRecords()
	blk.buildMerkleTree()
}

func (blk *Block) GetRecord(idx int) (record []byte, merklePath []int) {
	record = blk.records[idx]
	return
}

func (blk *Block) fillRecords() {
	res, _ := http.Get("https://sherlock-holm.es/stories/plain-text/bosc.txt")
	body, _ := ioutil.ReadAll(res.Body)
	sz := 512
	numRec := len(body) / sz // deliberate integer division
	for i := 0; i < numRec; i++ {
		record := body[i * sz: (i + 1) * sz]
		blk.records = append(blk.records, record)
	}
}


func (blk *Block) buildMerkleTree() {
	// Laye up tiers starting with the bottom tier
	bottomRow := []int{}
	for record, _ := range(blk.records) {
		bottomRow = append(bottomRow, hashOf(record))
	}
	rows := [][]byte{}
	lower = makeEvenLength(bottomRow)
	rows = append(rows, lower)
	for {
		if len(lower) == 1 {
			break;
		}
		upper := blk.buildNextMerkleRowUp(lower)
		rows = append(rows, upper)
		lower = upper
	}
	// Traverse tiers downwards from top tier
	nRows = len(rows)
	for i := nRows - 1; i >= 0; i-- {
		row := rows[i]
		depth := i
	}

	bottom row is row[0]		depth is nrows -1
	second row up is row[1]         depth is nrows -2
	third row up is row[2]         depth is nrows -3

	highest row is row[nrows - 1]  depth is nrows -

}