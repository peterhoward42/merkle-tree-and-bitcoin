package main

import (
	"github.com/peterhoward42/merkle-tree-and-bitcoin/src/bitcoin"
)

func main() {
	remoteNode := bitcoin.NewFullBitcoinNode()
	localNode := bitcoin.NewSpvBitcoinNode(&remoteNode)

	blockOfInterest := 2
	recordToFetch := 42

	localNode.FetchAndValidateRecordFromRemote(blockOfInterest, recordToFetch)
}
