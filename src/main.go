package main

import (
	"github.com/peterhoward42/merkle-tree-and-bitcoin/src/bitcoin"
)

func main() {
	remoteNode := bitcoin.NewFullBitcoinNode()
	localNode := bitcoin.NewSpvBitcoinNode(remoteNode)

	trustedMerkleRoot := remoteNode.MerkleRootForBlock()

	localNode.GetAndValidateRecord42FromRemoteNode(trustedMerkleRoot)
}
