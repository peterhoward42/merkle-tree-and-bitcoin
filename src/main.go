package main

import (
	"github.com/peterhoward42/merkle-tree-and-bitcoin/src/bitcoin"
)

func main() {
	remoteNode := bitcoin.NewFullBitcoinNode()
	localNode := bitcoin.NewSpvBitcoinNode(&remoteNode)

	localNode.GetAndValidateRecord42FromRemoteNode() // Prints OK
}
