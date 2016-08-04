package main

import (
    "fmt"
	"github.com/peterhoward42/merkle-tree-and-bitcoin/src/bitcoin"
)

func main() {
	remoteNode := bitcoin.NewFullBitcoinNode()
	localNode := bitcoin.NewSpvBitcoinNode(&remoteNode)

	blockOfInterest := 2
	recordToFetch := 42

	record, err := localNode.FetchAndValidateRecordFromRemote(
            blockOfInterest, recordToFetch)

    if err != nil {
        fmt.Printf("Fetch failed with: %v", err)
    } else {
        fmt.Printf("Fetched record of size %d bytes, and passed validation.",
            len(*record))
    }
}
